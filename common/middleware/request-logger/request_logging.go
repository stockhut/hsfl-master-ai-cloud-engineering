package request_logger

import (
	"log"
	"net/http"
	"time"
)

// responseWriterWithStatus wraps a http.ResponseWriter and records the first http status written
type responseWriterWithStatus struct {
	responseWriter http.ResponseWriter
	statusCode     *int
}

func (w *responseWriterWithStatus) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *responseWriterWithStatus) Write(bytes []byte) (int, error) {
	return w.responseWriter.Write(bytes)
}

func (w *responseWriterWithStatus) WriteHeader(statusCode int) {
	w.responseWriter.WriteHeader(statusCode)
	if w.statusCode == nil {
		w.statusCode = &statusCode
	}
}

func New(logger *log.Logger) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			startTime := time.Now()

			// we need a custom http.ResponseWriter so we can read the http status code
			responseWriter := responseWriterWithStatus{
				responseWriter: w,
			}

			next(&responseWriter, r)

			method := r.Method
			url := r.URL
			duration := time.Now().Sub(startTime)
			statusCode := responseWriter.statusCode

			// [GET] /foo/bar 404 12ms
			logger.Printf("[%s] %s %d %s", method, url, *statusCode, duration)

		}
	}
}
