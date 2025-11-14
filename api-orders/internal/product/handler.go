package product

import (
	"api-orders/configs"
	"api-orders/internal/middleware"
	"api-orders/internal/models"
	"api-orders/internal/user"
	"api-orders/pkg/request"
	"api-orders/pkg/response"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	Repository     *ProductRepository
	UserRepository *user.UserRepository
	AuthConfig     *configs.AuthConfig
}
type ProductHandler struct {
	Repository *ProductRepository
}

func NewHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		Repository: deps.Repository,
	}
	router.Handle("GET /products", middleware.IsAuthed(handler.get(), deps.AuthConfig, deps.UserRepository))
	router.Handle("POST /products", middleware.IsAuthed(handler.create(), deps.AuthConfig, deps.UserRepository))
	router.Handle("GET /products/{id}", middleware.IsAuthed(handler.getById(), deps.AuthConfig, deps.UserRepository))
	router.Handle("PATCH /products/{id}", middleware.IsAuthed(handler.update(), deps.AuthConfig, deps.UserRepository))
	router.Handle("DELETE /products/{id}", middleware.IsAuthed(handler.delete(), deps.AuthConfig, deps.UserRepository))
}
func (handler *ProductHandler) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Product get handler")
		page, err := request.PrepareParam[uint](&w, r, "query", "page", false)
		if err != nil {
			return
		}
		pageSize, err := request.PrepareParam[uint](&w, r, "query", "pageSize", false)
		if err != nil {
			return
		}
		products := handler.Repository.GetAll(page, pageSize)
		response.Json(w, products, http.StatusOK)
	}
}
func (handler *ProductHandler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Product getById handler")
		id, err := request.PrepareParam[uint](&w, r, "path", "id", true)
		if err != nil {
			return
		}
		product, err := handler.Repository.GetById(id)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusNotFound)
			return
		}
		response.Json(w, product, http.StatusOK)
	}
}
func (handler *ProductHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Product create handler")
		body, err := request.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}

		product := &models.Product{
			Name:        body.Name,
			Description: body.Description,
			Price:       body.Price,
			Quantity:    body.Quantity,
		}
		createdProduct, err := handler.Repository.Create(product)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdProduct, http.StatusCreated)
	}
}
func (handler *ProductHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Product update handler")
		body, err := request.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		id, err := request.PrepareParam[uint](&w, r, "path", "id", true)
		if err != nil {
			return
		}

		product, err := handler.Repository.Update(&models.Product{
			Model:       gorm.Model{ID: id},
			Name:        body.Name,
			Description: body.Description,
			Price:       body.Price,
			Quantity:    body.Quantity,
		})
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, product, http.StatusOK)
	}
}
func (handler *ProductHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Product delete handler")
		id, err := request.PrepareParam[uint](&w, r, "path", "id", true)
		if err != nil {
			return
		}

		_, err = handler.Repository.GetById(id)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.Repository.Delete(id)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, nil, http.StatusOK)
	}
}
