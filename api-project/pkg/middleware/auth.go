package middleware

import (
	"api-project/configs"
	"api-project/pkg/jwt"
	"api-project/pkg/response"
	"context"
	"log"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "contextEmailKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	response.JsonError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func IsAuthed(next http.Handler, config *configs.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeUnauthorized(w)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.New(config.JwtSecret).Parse(token)

		log.Println("Token", token)
		log.Println("Token is valid", isValid)
		log.Println("Token data", data)

		if !isValid {
			writeUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
