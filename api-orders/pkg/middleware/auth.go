package middleware

import (
	"api-orders/configs"
	"api-orders/pkg/jwt"
	"api-orders/pkg/response"
	"context"
	"log"
	"net/http"
	"strings"
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

		ctx := context.WithValue(r.Context(), jwt.ContextJwtKey, data)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
