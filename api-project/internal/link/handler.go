package link

import (
	"api-project/configs"
	"api-project/pkg/middleware"
	"api-project/pkg/request"
	"api-project/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	Repository *LinkRepository
	Config     *configs.AuthConfig
}
type LinkHandler struct {
	Repository *LinkRepository
}

func NewHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		Repository: deps.Repository,
	}
	router.HandleFunc("POST /links", handler.create())
	router.Handle("PATCH /links/{id}", middleware.IsAuthed(handler.update(), deps.Config))
	router.HandleFunc("DELETE /links/{id}", handler.delete())
	router.HandleFunc("GET /{hash}", handler.goTo())
}
func (handler *LinkHandler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.Repository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
func (handler *LinkHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		for {
			existsLink, _ := handler.Repository.GetByHash(link.Hash)
			if existsLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := handler.Repository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdLink, http.StatusCreated)
	}
}
func (handler *LinkHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		id, err := request.PrepareParam[uint](&w, r, "path", "id", true)
		if err != nil {
			return
		}

		contextEmail, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println("update handler contextEmail:", contextEmail)
		} else {
			fmt.Println("update handler unknown contextEmail")
		}

		link, err := handler.Repository.Update(&Link{
			Model: gorm.Model{ID: id},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, link, http.StatusOK)
	}
}
func (handler *LinkHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.Repository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.Repository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, nil, http.StatusOK)
	}
}
