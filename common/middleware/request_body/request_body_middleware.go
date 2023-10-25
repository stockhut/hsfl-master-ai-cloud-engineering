package request_body_middleware

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type BodyMiddlewareContextKeyType string

const BodyMiddlewareContextKey BodyMiddlewareContextKeyType = "body"

func GetBody[T any](ctx context.Context) T {
	return ctx.Value(BodyMiddlewareContextKey).(T)
}

func Body[T any](next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Printf("Failed to read body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var requestBody T
		if err := json.Unmarshal(body, &requestBody); err != nil {
			log.Printf("Failed to unmarshal body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), BodyMiddlewareContextKey, requestBody)
		next(w, r.WithContext(ctx))
	}

}
