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
	var repo inMemoryAccountRepository = inMemoryAccountRepository{
		accounts: make([]account, 0),
	}
	repo.accounts = append(repo.accounts, account{name: "Nele", email: "nele@nele.de", password: "xyz123"})
	repo.accounts = append(repo.accounts, account{name: "Alex", email: "alex@nele.de", password: "abc123"})
	repo.accounts = append(repo.accounts, account{name: "Fabi", email: "fabi@nele.de", password: "def123"})

	c := AccountController{accountRepo: &repo}

	router := router.New()
	router.POST("/account", c.handleCreateAccount)

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
