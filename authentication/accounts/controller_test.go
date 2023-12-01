package accounts

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_accounts "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/repository"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPostAccount(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tokenGenerator := jwt_util.NewJwtTokenGeneratorWithKey(privateKey)

	gomockController := gomock.NewController(t)

	mockRepo := mock_accounts.NewMockAccountRepository(gomockController)
	mockRepo.EXPECT().CreateAccount(gomock.Any()).Return(nil).Times(1)
	mockRepo.EXPECT().CheckDuplicate(gomock.Any()).Return(repository.NO_DUPLICATES, nil).Times(1)
	c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator}

	router := router.New()
	router.POST("/account", c.HandleCreateAccount)
	router.POST("/login", c.HandleLogin)

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

	t.Run("LoginAccount", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if payload is nil", func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", nil)

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.Code)

			assert.Nil(t, FindCookie(w.Result().Cookies(), "jwt"), "did not expect jwt token as cookie")
		})

		t.Run("should return 400 BAD REQUEST if password is wrong", func(t *testing.T) {
			testBody := `{"name":"Nele", "password":"wrong"}`

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(testBody))

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusNotFound, w.Code)

			assert.Nil(t, FindCookie(w.Result().Cookies(), "jwt"), "did not expect jwt token as cookie")
		})

		t.Run("should return 400 BAD REQUEST if username does not exist", func(t *testing.T) {
			testBody := `{"name":"doesnotexist", "password":"xyz123"}`

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(testBody))

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusNotFound, w.Code)

			assert.Nil(t, FindCookie(w.Result().Cookies(), "jwt"), "did not expect jwt token as cookie")
		})

		t.Run("should return 200 OK if login is successful", func(t *testing.T) {
			testBody := `{"name":"Nele", "password":"xyz123"}`

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(testBody))

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusOK, w.Code)

			assert.NotNil(t, FindCookie(w.Result().Cookies(), "jwt"), "expected jwt token as cookie")
		})
	})
}

func FindCookie(cookies []*http.Cookie, name string) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}
