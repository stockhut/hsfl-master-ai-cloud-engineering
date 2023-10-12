package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
	"github.com/stretchr/testify/assert"
)

func TestPostAccount(t *testing.T) {
	router := router.New()
	router.POST("/account", handleCreateAccount)

	t.Run("PostAccount", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if payload is nil", func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/account", nil)

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
		t.Run("should return 400 BAD REQUEST if payload is already existing", func(t *testing.T) {
			// TODO
		})

		t.Run("should return 201 CREATED when payload is valid", func(t *testing.T) {
			testBody := `{"name":"Bob","email": "bob@nele.de","password": "1234"}`

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/account", strings.NewReader(testBody))

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusCreated, w.Code)
		})
	})
}
