package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
	"github.com/stretchr/testify/assert"
)

func TestJwtMiddleware(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	publicKey := privateKey.Public()

	tokenGenerator := jwt_util.NewJwtTokenGeneratorWithKey(privateKey)

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		token := r.Context().Value("jwt").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		name, ok := claims["name"]

		if ok {
			fmt.Fprintf(w, "Hello %s", name)
		}

	}
	jwtMiddleware := ValidateJwtMiddleware(publicKey, testHandler)

	router := router.New()
	router.GET("/test", jwtMiddleware)

	t.Run("Jwt middleware", func(t *testing.T) {
		t.Run("should call next handler when jwt is valid", func(t *testing.T) {

			token, err := tokenGenerator.CreateToken(map[string]interface{}{
				"name": "testname",
			})
			if err != nil {
				panic(err)
			}

			w := httptest.NewRecorder()

			r := httptest.NewRequest(http.MethodGet, "/test", nil)
			r.AddCookie(&http.Cookie{
				Name:  "jwt",
				Value: token,
			})

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusOK, w.Code)

			body, err := io.ReadAll(w.Result().Body)
			if err != nil {
				panic(err)
			}

			fmt.Printf("body: %v\n", body)
			assert.Equal(t, []byte("Hello testname"), body)
		})

		t.Run("should return 401 UNAUTHORIZED when signature is invalid", func(t *testing.T) {

			token := "xyz123"

			w := httptest.NewRecorder()

			r := httptest.NewRequest(http.MethodGet, "/test", nil)
			r.AddCookie(&http.Cookie{
				Name:  "jwt",
				Value: token,
			})

			router.ServeHTTP(w, r)
			assert.Equal(t, http.StatusUnauthorized, w.Code)

			body, err := io.ReadAll(w.Result().Body)
			if err != nil {
				panic(err)
			}

			assert.Equal(t, 0, len(body))
		})

	})
}
