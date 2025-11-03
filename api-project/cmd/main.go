package main

import (
	"api-project/configs"
	"api-project/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.Load()
	router := http.NewServeMux()
	auth.New(router, auth.AuthHandlerDeps{
		Config: &conf.Auth,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
