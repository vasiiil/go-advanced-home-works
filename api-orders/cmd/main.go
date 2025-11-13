package main

import (
	"api-orders/configs"
	"api-orders/internal/product"
	"api-orders/pkg/db"
	"api-orders/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.Load()
	database := db.New(&conf.Db)
	router := http.NewServeMux()

	// #region Repositories
	productRepository := product.NewRepository(database)
	// #endregion Repositories

	// #region Handlers
	product.NewHandler(router, product.ProductHandlerDeps{
		Repository: productRepository,
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
