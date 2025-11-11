package link

import (
	"api-project/pkg/request"
	"api-project/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	Repository *LinkRepository
}
type LinkHandler struct {
	Repository *LinkRepository
}

func NewHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		Repository: deps.Repository,
	}
	router.HandleFunc("POST /links", handler.create())
	router.HandleFunc("PATCH /links/{id}", handler.update())
	router.HandleFunc("DELETE /links/{id}", handler.delete())
	router.HandleFunc("GET /{hash}", handler.goTo())
}
func (handler *LinkHandler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Link goTo handler")
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
		fmt.Println("Link create handler")
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
		fmt.Println("Link update handler")
		body, err := request.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := handler.Repository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
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
		fmt.Println("Link delete handler")
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
