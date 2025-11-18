package main

import (
	"api-project/configs"
	"api-project/internal/auth"
	"api-project/internal/link"
	"api-project/internal/stat"
	"api-project/internal/user"
	"api-project/internal/verify"
	"api-project/pkg/db"
	"api-project/pkg/email"
	"api-project/pkg/event"
	"api-project/pkg/middleware"
	"fmt"
	"net/http"
)

func App() http.Handler {
	conf := configs.Load()
	database := db.New(&conf.Db)
	router := http.NewServeMux()
	emailSender := email.New(&conf.Email)
	eventBus := event.NewEventBus()

	// #region Repositories
	linkRepository := link.NewRepository(database)
	userRepository := user.NewRepository(database)
	statRepository := stat.NewRepository(database)
	// #endregion Repositories

	// #region Services
	authService := auth.NewService(userRepository)
	statService := stat.NewService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	// #endregion Services

	go statService.AddClick()

	// #region Handlers
	auth.NewHandler(router, auth.AuthHandlerDeps{
		Config:  &conf.Auth,
		Service: authService,
	})
	link.NewHandler(router, link.LinkHandlerDeps{
		Config:     &conf.Auth,
		Repository: linkRepository,
		EventBus:   eventBus,
	})
	stat.NewHandler(router, stat.StatHandlerDeps{
		Config:     &conf.Auth,
		Repository: statRepository,
	})
	// #endregion Handlers

	// #region Middlewares
	stackMiddleware := middleware.Chain(
		middleware.Cors,
		middleware.Logging,
	)
	// #endregion Middlewares

	verify.New(router, verify.EmailHandlerDeps{
		EmailSender: emailSender,
	})
	return stackMiddleware(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
