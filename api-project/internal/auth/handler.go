package auth

import (
	"api-project/configs"
	"api-project/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	Config *configs.AuthConfig
}
type AuthHandler struct {
	Config *configs.AuthConfig
}

func New(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}
func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login handler")
		fmt.Printf("Secret is: %v\n", handler.Config.Secret)
		resp := LoginResponse{
			Token: handler.Config.Secret,
		}
		response.Json(w, resp, http.StatusAccepted)
	}
}
func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register handler")
	}
}
