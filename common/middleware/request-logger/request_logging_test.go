package request_logger

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_responseWriterWithStatus_WriteHeader(t *testing.T) {

	t.Run("records the first written status code and calls wrapped methods", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		responseWriter := responseWriterWithStatus{
			responseWriter: recorder,
		}

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("foo", "bar")

			w.WriteHeader(http.StatusTeapot)
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, "test")
		}

		handler(&responseWriter, nil)

		assert.Equal(t, http.StatusTeapot, *responseWriter.statusCode)
		assert.Equal(t, recorder.Result().StatusCode, *responseWriter.statusCode)

		body, err := io.ReadAll(recorder.Body)

		assert.Nil(t, err)
		assert.Equal(t, "test", string(body))

		fooHeaders := recorder.Result().Header.Values("foo")
		assert.Len(t, fooHeaders, 1)
		assert.Equal(t, "bar", fooHeaders[0])
	})
}

func Test_request_logger_middleware(t *testing.T) {

	t.Run("logs something", func(t *testing.T) {

		recorder := httptest.NewRecorder()

		// record bytes written by the middleware
		logData := make([]byte, 0)
		logWriter := bytes.NewBuffer(logData)

		logger := log.New(logWriter, "", 0)

		handlerCalled := false

		mw := New(logger)

		r := httptest.NewRequest(http.MethodPost, "/test", nil)
		handler := mw(func(w http.ResponseWriter, r *http.Request) {
			handlerCalled = true
			w.WriteHeader(http.StatusTeapot)
		})

		handler(recorder, r)

		assert.NotEmpty(t, true, logData)
		assert.Equal(t, true, handlerCalled)

		logString := logWriter.String()
		assert.Contains(t, logString, "POST")
		assert.Contains(t, logString, "418")
	})

	t.Run("recovers from panic in subsequent handlers", func(t *testing.T) {

		recorder := httptest.NewRecorder()

		// record bytes written by the middleware
		logData := make([]byte, 0)
		logWriter := bytes.NewBuffer(logData)

		logger := log.New(logWriter, "", 0)
		mw := New(logger)

		r := httptest.NewRequest(http.MethodPost, "/test", nil)
		handler := mw(func(w http.ResponseWriter, r *http.Request) {
			panic("test")
		})

		handler(recorder, r)

		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)

		logString := logWriter.String()
		assert.Contains(t, logString, "PANIC")
	})

	t.Run("should not break when the handler did not set a response code", func(t *testing.T) {

		recorder := httptest.NewRecorder()

		// record bytes written by the middleware
		logData := make([]byte, 0)
		logWriter := bytes.NewBuffer(logData)

		logger := log.New(logWriter, "", 0)
		mw := New(logger)

		r := httptest.NewRequest(http.MethodPost, "/test", nil)
		handler := mw(func(w http.ResponseWriter, r *http.Request) {

		})

		handler(recorder, r)

		//assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)

		//logString := logWriter.String()
		//assert.Contains(t, logString, "PANIC")
	})
}
