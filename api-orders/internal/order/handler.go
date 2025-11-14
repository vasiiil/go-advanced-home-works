package order

import (
	"api-orders/configs"
	"api-orders/internal/middleware"
	"api-orders/internal/user"
	"api-orders/pkg/request"
	"api-orders/pkg/response"
	"fmt"
	"net/http"
)

type OrderHandlerDeps struct {
	Repository     *OrderRepository
	UserRepository *user.UserRepository
	AuthConfig     *configs.AuthConfig
}
type OrderHandler struct {
	Repository *OrderRepository
}

func NewHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		Repository: deps.Repository,
	}
	router.Handle("POST /orders", middleware.IsAuthed(handler.create(), deps.AuthConfig, deps.UserRepository))
	router.Handle("GET /orders/{id}", middleware.IsAuthed(handler.getById(), deps.AuthConfig, deps.UserRepository))
	router.Handle("GET /my-orders", middleware.IsAuthed(handler.getUserOrders(), deps.AuthConfig, deps.UserRepository))
}
func (handler *OrderHandler) getUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Order getUserOrders handler")
		page, err := request.PrepareParam[uint](&w, r, "query", "page", false)
		if err != nil {
			return
		}
		pageSize, err := request.PrepareParam[uint](&w, r, "query", "pageSize", false)
		if err != nil {
			return
		}

		ctxUser := middleware.GetUserFromContext(w, r)
		if ctxUser == nil {
			return
		}

		orders, err := handler.Repository.GetByUserId(ctxUser.ID, page, pageSize)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, orders, http.StatusOK)
	}
}
func (handler *OrderHandler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Order getById handler")
		id, err := request.PrepareParam[uint](&w, r, "path", "id", true)
		if err != nil {
			return
		}
		ctxUser := middleware.GetUserFromContext(w, r)
		if ctxUser == nil {
			return
		}
		order, err := handler.Repository.GetById(ctxUser.ID, id)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusNotFound)
			return
		}
		response.Json(w, order, http.StatusOK)
	}
}
func (handler *OrderHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Order create handler")
		body, err := request.HandleBody[OrderCreateRequest](&w, r)
		if err != nil {
			return
		}
		ctxUser := middleware.GetUserFromContext(w, r)
		if ctxUser == nil {
			return
		}
		createdOrder, err := handler.Repository.Create(ctxUser.ID, body.ProductIds)
		if err != nil {
			response.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdOrder, http.StatusCreated)
	}
}
