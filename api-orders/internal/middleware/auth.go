package middleware

import (
	"api-orders/configs"
	"api-orders/internal/models"
	"api-orders/internal/user"
	"api-orders/pkg/jwt"
	"api-orders/pkg/response"
	"context"
	"log"
	"net/http"
	"strings"
)

type TContextUserKey string

const (
	ContextUserKey TContextUserKey = "contextUserKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	response.JsonError(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func IsAuthed(next http.Handler, config *configs.AuthConfig, userRepo *user.UserRepository) http.Handler {
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

		_user, err := userRepo.GetBySessionId(data.SessionId)
		if err != nil {
			writeUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, *_user)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func GetUserFromContext(w http.ResponseWriter, r *http.Request) *models.User {
	_user, ok := r.Context().Value(ContextUserKey).(models.User)
	if !ok {
		writeUnauthorized(w)
		return nil
	}
	return &_user
}
