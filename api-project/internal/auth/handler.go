package auth

import (
	"api-project/configs"
	"api-project/pkg/request"
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
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println("Login handler")
		fmt.Printf("Secret is: %v\n", handler.Config.Secret)
		fmt.Println("Payload:", payload)
		resp := LoginResponse{
			Token: handler.Config.Secret,
		}
		response.Json(w, resp, http.StatusAccepted)
	}
}
func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println("Register handler")
		fmt.Println("Payload:", payload)
	}
}
