package verify

import (
	"api-project/configs"
	"fmt"
	"net/http"
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
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Email send handler")
		w.WriteHeader(http.StatusOK)
	}
}
func (handler *EmailHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Email verify handler")
		w.WriteHeader(http.StatusOK)
	}
}
