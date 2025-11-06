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

func New(router *http.ServeMux, deps EmailHandlerDeps) {
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
		
		//TODO здесь будет сохранение хеша и емейла
		
		text := fmt.Sprintf("<h1>Hello, %s!</h1><p>Please, verify your email %s, following by <a href=\"http://localhost/verify/%s\">link</a></p>", payload.Name, payload.Email, hash)
		mailer := email.New(handler.Config)
		err = mailer.Send(handler.Config.Email, "Verifying email", text)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
func (handler *EmailHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// http://localhost/verify/0d04af71-4b15-4211-bb44-eff52b3dbf29
		fmt.Println("Email verify handler")
		hash := r.PathValue("hash")
		if hash == "" {
			response.BadRequestJson(w, "Empty Hash")
			return
		}

		//todo здесь ищем запись с таким хэшем
		// пока что, пока нет базы данных, захардкодим
		if hash == "0d04af71-4b15-4211-bb44-eff52b3dbf29" {
			response.Json(w, VerifyResponse{
				Success: true,
			}, http.StatusOK)
		} else {
			response.Json(w, VerifyResponse{
				Success: false,
			}, http.StatusUnauthorized)
			//todo здесь будет удаление
		}
	}
}
