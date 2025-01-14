package accounts

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	mock_pwhash "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks/pwhash_mocks"
	mock_accounts "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks/repository_mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountController(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tokenGenerator := jwt_util.NewJwtTokenGeneratorWithKey(privateKey)

	t.Run("CreateAccount", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if payload is nil", func(t *testing.T) {
			gomockController := gomock.NewController(t)

			mockRepo := mock_accounts.NewMockRepository(gomockController)

			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/account", nil)

			c.HandleCreateAccount(w, r)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
		//AB HIER no tests to run
		t.Run("should return 400 BAD REQUEST if payload is already existing", func(t *testing.T) {
			testBody := `{"name":"Bob","email": "bob@nele.de","password": "1234"}`
			gomockController := gomock.NewController(t)

			mockRepo := mock_accounts.NewMockRepository(gomockController)
			mockRepo.EXPECT().CheckDuplicate(gomock.Any(), gomock.Any()).Return(ErrDuplicateEmail).Times(1)

			mockPwHasher := mock_pwhash.NewMockPasswordHasher(gomockController)
			mockPwHasher.EXPECT().Hash("1234").Return([]byte("passwordhash"), nil).Times(1)

			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator, pwHasher: mockPwHasher}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/account", strings.NewReader(testBody))

			c.HandleCreateAccount(w, r)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("should return 201 CREATED when payload is valid", func(t *testing.T) {

			testBody := `{"name":"Bob","email": "bob@nele.de","password": "1234"}`
			modelAccount := model.Account{
				Name:         "Bob",
				Email:        "bob@nele.de",
				PasswordHash: []byte("passwordhash"),
			}
			gomockController := gomock.NewController(t)

			mockRepo := mock_accounts.NewMockRepository(gomockController)
			mockRepo.EXPECT().CheckDuplicate(gomock.Any(), modelAccount).Return(nil).Times(1)
			mockRepo.EXPECT().CreateAccount(gomock.Any(), modelAccount).Return(nil).Times(1)

			mockPwHasher := mock_pwhash.NewMockPasswordHasher(gomockController)
			mockPwHasher.EXPECT().Hash("1234").Return([]byte("passwordhash"), nil).Times(1)

			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator, pwHasher: mockPwHasher}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/account", strings.NewReader(testBody))

			c.HandleCreateAccount(w, r)

			assert.Equal(t, http.StatusCreated, w.Code)
		})

		t.Run("should return 500 INTERNAL SERVER ERROR if password hash fails", func(t *testing.T) {

			testBody := `{"name":"Bob","email": "bob@nele.de","password": "1234"}`

			gomockController := gomock.NewController(t)

			mockRepo := mock_accounts.NewMockRepository(gomockController)

			mockPwHasher := mock_pwhash.NewMockPasswordHasher(gomockController)
			mockPwHasher.EXPECT().Hash("1234").Return([]byte{}, errors.New("error")).Times(1)

			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator, pwHasher: mockPwHasher}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/account", strings.NewReader(testBody))

			c.HandleCreateAccount(w, r)
			assert.Equal(t, http.StatusInternalServerError, w.Code)

			assert.Nil(t, FindCookie(w.Result().Cookies(), "jwt"), "did not expect jwt token as cookie")
		})
	})

	t.Run("LoginAccount", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if payload is nil", func(t *testing.T) {
			gomockController := gomock.NewController(t)

			mockRepo := mock_accounts.NewMockRepository(gomockController)

			mockPwHasher := mock_pwhash.NewMockPasswordHasher(gomockController)

			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator, pwHasher: mockPwHasher}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", nil)

			c.HandleLogin(w, r)
			assert.Equal(t, http.StatusBadRequest, w.Code)

			assert.Nil(t, FindCookie(w.Result().Cookies(), "jwt"), "did not expect jwt token as cookie")
		})

		t.Run("should return 400 BAD REQUEST if password is wrong", func(t *testing.T) {
			gomockController := gomock.NewController(t)
			modelAccount := model.Account{
				Name:         "Nele",
				Email:        "nele@nele.de",
				PasswordHash: []byte("storedhash"),
			}

			mockRepo := mock_accounts.NewMockRepository(gomockController)
			mockRepo.EXPECT().FindAccount(gomock.Any(), modelAccount.Name).Return(&modelAccount, nil).Times(1)

			mockPwHasher := mock_pwhash.NewMockPasswordHasher(gomockController)
			mockPwHasher.EXPECT().Verify([]byte("storedhash"), "wrong").Return(false).Times(1)

			testBody := `{"name":"Nele", "password":"wrong"}`
			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator, pwHasher: mockPwHasher}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(testBody))

			c.HandleLogin(w, r)
			assert.Equal(t, http.StatusNotFound, w.Code)
			assert.Nil(t, FindCookie(w.Result().Cookies(), "jwt"), "did not expect jwt token as cookie")
		})

		t.Run("should return 400 BAD REQUEST if username does not exist", func(t *testing.T) {
			gomockController := gomock.NewController(t)
			mockRepo := mock_accounts.NewMockRepository(gomockController)
			mockRepo.EXPECT().FindAccount(gomock.Any(), "doesnotexist").Return(nil, nil).Times(1)

			mockPwHasher := mock_pwhash.NewMockPasswordHasher(gomockController)

			testBody := `{"name":"doesnotexist", "password":"xyz123"}`
			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator, pwHasher: mockPwHasher}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(testBody))

			c.HandleLogin(w, r)
			assert.Equal(t, http.StatusNotFound, w.Code)
			assert.Nil(t, FindCookie(w.Result().Cookies(), "jwt"), "did not expect jwt token as cookie")
		})

		t.Run("should return 200 OK if login is successful", func(t *testing.T) {
			gomockController := gomock.NewController(t)
			modelAccount := model.Account{
				Name:         "Nele",
				Email:        "nele@nele.de",
				PasswordHash: []byte("storedhash"),
			}

			mockRepo := mock_accounts.NewMockRepository(gomockController)
			mockRepo.EXPECT().FindAccount(gomock.Any(), modelAccount.Name).Return(&modelAccount, nil).Times(1)

			mockPwHasher := mock_pwhash.NewMockPasswordHasher(gomockController)
			mockPwHasher.EXPECT().Verify([]byte("storedhash"), "1234").Return(true).Times(1)

			testBody := `{"name":"Nele", "password":"1234"}`
			c := Controller{accountRepo: mockRepo, tokenGenerator: *tokenGenerator, pwHasher: mockPwHasher}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(testBody))

			c.HandleLogin(w, r)
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
