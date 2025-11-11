package main

import (
	"api-orders/configs"
	"api-orders/internal/product"
	"api-orders/pkg/db"
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

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
