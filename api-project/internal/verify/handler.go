package verify

import (
	"api-project/configs"
	"api-project/pkg/email"
	"api-project/pkg/request"
	"api-project/pkg/response"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type EmailHandlerDeps struct {
	Config *configs.EmailConfig
}
type EmailHandler struct {
	Config *configs.EmailConfig
}

var verifyEmailsMap map[string]string

func New(router *http.ServeMux, deps EmailHandlerDeps) {
	verifyEmailsMap = make(map[string]string)
	handler := &EmailHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.send())
	router.HandleFunc("GET /verify/{hash}", handler.verify())
}
func (handler *EmailHandler) send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Email send handler")
		payload, err := request.HandleBody[VerifyRequest](&w, r)
		if err != nil {
			return
		}

		hash := uuid.New().String()
		verifyEmailsMap[hash] = payload.Email

		text := fmt.Sprintf("<h1>Hello, %s!</h1><p>Please, verify your email %s, following by <a href=\"http://localhost/verify/%s\">link</a></p>", payload.Name, payload.Email, hash)
		mailer := email.New(handler.Config)
		err = mailer.Send(handler.Config.Email, "Verifying email", text)
		fmt.Println("Hash:", hash)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
func (handler *EmailHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Email verify handler")
		hash := r.PathValue("hash")
		if hash == "" {
			response.BadRequestJson(w, "Empty Hash")
			return
		}

		email, ok := verifyEmailsMap[hash]
		if ok {
			fmt.Printf("Email %s is verified\n", email)
			delete(verifyEmailsMap, hash)
			response.Json(w, VerifyResponse{
				Success: true,
			}, http.StatusOK)
		} else {
			response.Json(w, VerifyResponse{
				Success: false,
			}, http.StatusUnauthorized)
		}
	}
}
