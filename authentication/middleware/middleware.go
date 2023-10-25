package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type JwtContextKeyType string

const JwtContextKey JwtContextKeyType = "jwt-claim"

func ValidateJwtMiddleware(publicKey any) func(http.HandlerFunc) http.HandlerFunc {

	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("jwt")
			if err != nil {
				fmt.Println("no jwt cookie")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
				fmt.Println(t)
				return publicKey, nil
			})

			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)

				return
			}
			claims := token.Claims.(jwt.MapClaims)
			ctx := context.WithValue(r.Context(), JwtContextKey, claims)

			next(w, r.WithContext(ctx))
		}
	}
}
