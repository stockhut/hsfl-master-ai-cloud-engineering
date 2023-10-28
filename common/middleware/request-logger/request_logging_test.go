package request_logger

import (
	"github.com/stretchr/testify/assert"
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
