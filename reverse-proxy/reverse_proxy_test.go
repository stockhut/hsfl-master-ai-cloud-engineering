package reverse_proxy

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
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

	forward(recorder, request, service)

	assert.Equal(t, http.StatusTeapot, recorder.Code, "Forwarded wrong http status code")

	body, err := io.ReadAll(recorder.Result().Body)
	assert.Nil(t, err)

	assert.Equal(t, "request body", requestBody)
	assert.Equal(t, "response body", string(body))

	assert.Equal(t, "TestHeaderValue", recorder.Result().Header.Get("TestHeader"))
}
