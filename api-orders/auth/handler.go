package auth

import (
	"api-orders/configs"
	"api-orders/pkg/jwt"
	"api-orders/pkg/request"
	"api-orders/pkg/response"
	"fmt"
	"net/http"
	"regexp"
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
	router.HandleFunc("POST /auth/verify", handler.verifyCode())
}
func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login handler")
		handler.Service.PrintSessions("before login")
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		matched, err := regexp.MatchString("^(7|8)?9\\d{9}$", body.Phone)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !matched {
			response.BadRequestJson(w, "invalid phone format")
			return
		}

		sessionId, err := handler.Service.Login(body.Phone)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		handler.Service.PrintSessions("after login")
		response.Json(w, LoginResponse{
			SessionId: sessionId,
		}, http.StatusOK)
	}
}
func (handler *AuthHandler) verifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify Code handler")
		handler.Service.PrintSessions("before verify code")
		body, err := request.HandleBody[VerifyCodeRequest](&w, r)
		if err != nil {
			return
		}
		verified := handler.Service.VerifyCode(body.SessionId, body.Code)
		if !verified {
			response.JsonError(w, ErrWrongVerificationCode.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.New(handler.Config.JwtSecret).Create(body.SessionId)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, VerifyCodeResponse{
			Token: token,
		}, http.StatusOK)
	}
}
