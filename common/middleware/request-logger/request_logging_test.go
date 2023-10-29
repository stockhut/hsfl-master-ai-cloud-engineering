package request_logger

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_responseWriterWithStatus_WriteHeader(t *testing.T) {

	t.Run("records the first written status code", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		responseWriter := responseWriterWithStatus{
			responseWriter: recorder,
		}

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
			w.WriteHeader(http.StatusOK)
		}

		handler(&responseWriter, nil)

		assert.Equal(t, http.StatusTeapot, *responseWriter.statusCode)
		assert.Equal(t, recorder.Result().StatusCode, *responseWriter.statusCode)
	})
}

func Test_request_logger_middleware(t *testing.T) {

	t.Run("logs someething", func(t *testing.T) {

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
}
