package main

import (
	"api-project/configs"
	"api-project/internal/auth"
	"api-project/internal/link"
	"api-project/internal/verify"
	"api-project/pkg/db"
	"api-project/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.Load()
	database := db.New(&conf.Db)
	router := http.NewServeMux()

	// #region Repositories
	linkRepository := link.NewRepository(database)
	// #endregion Repositories

	// #region Handlers
	auth.NewHandler(router, auth.AuthHandlerDeps{
		Config: &conf.Auth,
	})
	link.NewHandler(router, link.LinkHandlerDeps{
		Repository: linkRepository,
	})
	// #endregion Handlers

	// Middlewares
	stackMiddleware := middleware.Chain(
		middleware.Cors,
		middleware.Logging,
	)

	verify.New(router, verify.EmailHandlerDeps{
		Config: &conf.Email,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: stackMiddleware(router),
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
