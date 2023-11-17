package reverse_proxy

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestForward(t *testing.T) {

	var requestBody string
	testHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		b, _ := io.ReadAll(r.Body)

		requestBody = string(b)

		w.Header().Set("TestHeader", "TestHeaderValue")
		w.WriteHeader(http.StatusTeapot)

		fmt.Fprint(w, "response body")

	}))

	defer testHttpServer.Close()

	fmt.Printf("test server listening on %s\n", testHttpServer.Listener.Addr())
	service := Service{
		Name:       "TestService",
		Route:      "/some/test/route",
		TargetHost: testHttpServer.Listener.Addr().String(),
	}

	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/some/test/route", bytes.NewReader([]byte("request body")))

	assert.Nil(t, err)

	Forward(recorder, request, service.TargetHost)

	assert.Equal(t, http.StatusTeapot, recorder.Code, "Forwarded wrong http status code")

	body, err := io.ReadAll(recorder.Result().Body)
	assert.Nil(t, err)

	assert.Equal(t, "request body", requestBody)
	assert.Equal(t, "response body", string(body))

	assert.Equal(t, "TestHeaderValue", recorder.Result().Header.Get("TestHeader"))
}

func TestPickService(t *testing.T) {

	t.Run("picks matching route", func(t *testing.T) {

		services := []Service{
			{
				Name:       "a",
				Route:      "/a",
				TargetHost: "a.example.org",
			},
			{
				Name:       "a",
				Route:      "/a/foo",
				TargetHost: "a.example.org",
			},
			{
				Name:       "b",
				Route:      "/b",
				TargetHost: "b.example.org",
			},
			{
				Name:       "wildcard",
				Route:      "/",
				TargetHost: "slash.example.org",
			},
		}

		for _, s := range services {

			t.Run(fmt.Sprintf("%s -> %s", s.Route, s.TargetHost), func(t *testing.T) {

				u := url.URL{
					Path: s.Route,
				}

				service := pickService(services, &u)

				assert.Equal(t, s.TargetHost, service.TargetHost)
			})
		}
	})

	t.Run("returns nil when no match is found", func(t *testing.T) {

		assert.Nil(t, pickService([]Service{}, &url.URL{Path: "/foo/bar"}))
	})
}

func TestReverseProxy_ServeHTTP(t *testing.T) {

	services := []Service{
		{
			Name:       "a",
			Route:      "/a",
			TargetHost: "a.example.org",
		},
	}

	t.Run("uses provided HttpForwarder", func(t *testing.T) {

		called := false
		fwd := func(w http.ResponseWriter, r *http.Request, host string) error {
			called = true
			return nil
		}

		rp := newWithForwarder(log.New(os.Stdout, "", 0), fwd, services)

		recorder := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "https://a.example.org/a", nil)

		assert.Nil(t, err, "http request failed")

		rp.ServeHTTP(recorder, req)

		assert.True(t, called)

	})

	t.Run("responds with 502 BAD GATEWAY when forwarding fails", func(t *testing.T) {

		fwd := func(w http.ResponseWriter, r *http.Request, host string) error {
			return errors.New("some error")
		}

		rp := newWithForwarder(log.New(os.Stdout, "", 0), fwd, services)

		recorder := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "https://a.example.org/a", nil)

		assert.Nil(t, err, "http request failed")

		rp.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})

	t.Run("responds with 502 BAD GATEWAY when no matching service is found", func(t *testing.T) {

		called := false
		fwd := func(w http.ResponseWriter, r *http.Request, host string) error {
			called = true
			return errors.New("some error")
		}

		rp := newWithForwarder(log.New(os.Stdout, "", 0), fwd, services)

		recorder := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "https://a.example.org/some/different/route", nil)

		assert.Nil(t, err, "http request failed")

		rp.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
		assert.False(t, called, "Proxy should not call the forward method when no service is found")
	})
}
