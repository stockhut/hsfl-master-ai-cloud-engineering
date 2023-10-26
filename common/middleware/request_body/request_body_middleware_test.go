package request_body

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBodyMiddleware(t *testing.T) {

	type testBodyStruct struct {
		Foo string
	}

	t.Run("BodyMiddleware", func(t *testing.T) {
		t.Run("should extract body", func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"Foo": "bar"}`))

			Body[testBodyStruct](func(w http.ResponseWriter, r *http.Request) {

				fmt.Printf("%v\n", r.Context())
				body := r.Context().Value(BodyMiddlewareContextKey).(testBodyStruct)

				assert.Equal(t, "bar", body.Foo)
				w.WriteHeader(http.StatusOK)
			})(w, r)

			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("should return 400 BAD REQUEST if body is invalid", func(t *testing.T) {

			tests := []string{``, `{`}

			for _, testBody := range tests {

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(testBody))

				Body[testBodyStruct](func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				})(w, r)

				assert.Equal(t, http.StatusBadRequest, w.Code)
			}

		})
	})
}
