package auth

import (
	"api-project/configs"
	"api-project/pkg/request"
	"api-project/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	Config  *configs.AuthConfig
	Service *AuthService
}
type AuthHandler struct {
	Config  *configs.AuthConfig
	Service *AuthService
}

func NewHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:  deps.Config,
		Service: deps.Service,
	}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}
func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login handler")
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		fmt.Printf("Secret is: %v\n", handler.Config.Secret)
		fmt.Println("Payload:", body)

		email, err := handler.Service.Login(body.Email, body.Password)
		if err != nil {
			response.BadRequestJson(w, err.Error())
			return
		}

		// resp := LoginResponse{
		// 	Token: handler.Config.Secret,
		// }
		response.Json(w, LoginResponse{
			Token: email,
		}, http.StatusAccepted)
	}
}
func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register handler")
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.Service.Register(body.Email, body.Password, body.Name)
		if err != nil {
			response.BadRequestJson(w, err.Error())
			return
		}

		response.Json(w, RegisterResponse{
			Token: email,
		}, http.StatusCreated)
	}
}
