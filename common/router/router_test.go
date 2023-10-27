package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("should return 404 NOT FOUND if path is unknown", func(t *testing.T) {
		// given
		router := New()
		router.GET("/the/route/without/params", func(w http.ResponseWriter, r *http.Request) {})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/unknown/route", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should route to correct handler without params", func(t *testing.T) {
		// given
		router := New()
		router.GET("/the/route/without/params", func(w http.ResponseWriter, r *http.Request) {})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/the/route/without/params", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should route to correct handler with params and pick correct method", func(t *testing.T) {
		// given

		var method string

		tests := []string{
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete,
		}

		router := New()
		var ctx context.Context

		handler := func(w http.ResponseWriter, r *http.Request) {
			ctx = r.Context()
			method = r.Method
		}
		router.GET("/the/:route/with/:params", handler)
		router.PUT("/the/:route/with/:params", handler)
		router.POST("/the/:route/with/:params", handler)
		router.DELETE("/the/:route/with/:params", handler)

		for _, testHttpMethod := range tests {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(testHttpMethod, "/the/route/with/params", nil)

			// when
			router.ServeHTTP(w, r)

			// then
			assert.Equal(t, testHttpMethod, method)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "route", ctx.Value("route"))
			assert.Equal(t, "params", ctx.Value("params"))

		}
	})
}
