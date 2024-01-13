package loadtest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpRequestBytesFromTarget(t *testing.T) {

	t.Run("Builds HTTP request with body and Content-Length", func(t *testing.T) {

		target := Target{
			Method: http.MethodPost,
			Path:   "/test",
			Body:   []byte("Hello World"),
		}

		handled := make(chan any)

		var requestBody string
		var contentLength int64
		var path string
		testHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			b, _ := io.ReadAll(r.Body)
			contentLength = r.ContentLength

			requestBody = string(b)
			path = r.URL.Path

			handled <- nil
		}))

		defer testHttpServer.Close()
		host := testHttpServer.Listener.Addr().String()

		conn, err := net.Dial("tcp", host)
		defer conn.Close()
		if err != nil {
			t.Fatal(err)
			return
		}

		err = MakeHttpRequest(conn, fmt.Sprintf("Host: %s", host), target)

		assert.Nil(t, err)

		select {
		case <-time.After(1 * time.Second):
			t.FailNow()
		case <-handled:
		}
		assert.Equal(t, "Hello World", requestBody)
		assert.Equal(t, int64(11), contentLength)
		assert.Equal(t, "/test", path)

	})

	t.Run("Builds HTTP request without body", func(t *testing.T) {

		target := Target{
			Method: http.MethodGet,
			Path:   "/test",
			Body:   make([]byte, 0),
		}

		handled := make(chan any)

		var path string
		testHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			path = r.URL.Path

			handled <- nil
		}))

		defer testHttpServer.Close()
		host := testHttpServer.Listener.Addr().String()

		conn, err := net.Dial("tcp", host)
		defer conn.Close()
		if err != nil {
			t.Fatal(err)
			return
		}

		err = MakeHttpRequest(conn, fmt.Sprintf("Host: %s", host), target)

		assert.Nil(t, err)

		select {
		case <-time.After(1 * time.Second):
			t.FailNow()
		case <-handled:
		}
		assert.Equal(t, "/test", path)

	})
}

func Test_rpsAtTime(t *testing.T) {

	phases := []Phase{
		{
			Rps:      100,
			Duration: 10 * time.Second,
			Rampup:   3 * time.Second,
		},
		{
			Rps:      200,
			Duration: 10 * time.Second,
			Rampup:   3 * time.Second,
		},
	}

	testCases := []struct {
		name string
		t    time.Duration
		rps  float64
	}{
		{
			name: "rampup phase 1",
			t:    1500 * time.Millisecond,
			rps:  50,
		},
		{
			name: "phase 1",
			t:    4 * time.Second,
			rps:  100,
		},
		{
			name: "rampup phase 2",
			t:    14500 * time.Millisecond,
			rps:  150,
		},
		{
			name: "phase 2",
			t:    18 * time.Second,
			rps:  200,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			rps := RpsAfterTime(phases, tc.t)

			assert.Equal(t, tc.rps, rps)
		})
	}

}

func Test_lerp(t *testing.T) {

	start := 0.0
	end := 100.0

	totalTime := 10 * time.Second
	elapsed := 5 * time.Second

	assert.Equal(t, 50.0, lerp(start, end, elapsed, totalTime))
}

func Test_httpStatusIsError(t *testing.T) {
	testCases := []struct {
		code    int
		isError bool
	}{
		{
			http.StatusOK, false,
		},
		{
			http.StatusAccepted, false,
		},
		{
			http.StatusBadRequest, true,
		},
		{
			http.StatusUnauthorized, true,
		},
		{
			http.StatusInternalServerError, true,
		},
		{
			http.StatusBadGateway, true,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, HttpStatusIsError(tc.code), tc.isError)
	}

}
