package request_logger

import (
	tc "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/termcolor"
	"log"
	"net/http"
	"strconv"
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

			method := httpMethodColorizer(r.Method)
			url := r.URL
			duration := time.Now().Sub(startTime)
			var statusCode string
			statusCode = statusCodeColorizer(*responseWriter.statusCode)

			// [GET] /foo/bar 404 12ms
			logger.Printf("[%s] %s %s %s", method, url, statusCode, duration)
		}
	}
}

func statusCodeColorizer(code int) string {
	return tc.Bg(httpStatusColor(code))(strconv.Itoa(code))
}

func httpStatusColor(code int) tc.BgColor {

	switch {
	case code < 200:
		return tc.BgBrightBlue // 100-199
	case code < 300:
		return tc.BgGreen // 200-299
	case code < 400:
		return tc.BgCyan // 300-299
	case code < 500:
		return tc.BgRed // 400-499
	case code < 600:
		return tc.BgRed // 500-599
	default:
		return tc.BgUndefined
	}
}

func httpMethodColorizer(method string) string {
	colors := map[string]tc.BgColor{
		http.MethodGet:    tc.BgGreen,
		http.MethodPost:   tc.BgBlue,
		http.MethodPut:    tc.BgCyan,
		http.MethodDelete: tc.BgRed,
	}

	return tc.Bg(colors[method])(method)
}
