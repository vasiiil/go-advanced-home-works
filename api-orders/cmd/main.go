package main

import (
	"api-orders/configs"
	"api-orders/internal/auth"
	"api-orders/internal/order"
	"api-orders/internal/product"
	"api-orders/internal/user"
	"api-orders/pkg/db"
	"api-orders/pkg/middleware"
	"api-orders/pkg/sms"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.Load()
	database := db.New(&conf.Db)
	router := http.NewServeMux()
	smsSender := sms.New(&conf.Sms)

	// #region Repositories
	productRepository := product.NewRepository(database)
	userRepository := user.NewRepository(database)
	orderRepository := order.NewRepository(database)
	// #endregion Repositories

	// #region Services
	authService := auth.NewService(auth.AuthServiceDeps{
		SmsSender:      smsSender,
		UserRepository: userRepository,
	})
	// #endregion Services

	// #region Handlers
	auth.NewHandler(router, auth.AuthHandlerDeps{
		Config:  &conf.Auth,
		Service: authService,
	})
	product.NewHandler(router, product.ProductHandlerDeps{
		Repository:     productRepository,
		UserRepository: userRepository,
		AuthConfig:     &conf.Auth,
	})
	order.NewHandler(router, order.OrderHandlerDeps{
		Repository:     orderRepository,
		UserRepository: userRepository,
		AuthConfig:     &conf.Auth,
	})
	// #endregion Handlers

	// Middlewares
	stackMiddleware := middleware.Chain(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stackMiddleware(router),
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
