package main

import (
	"api-orders/configs"
	"api-orders/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.Load()
	_ = db.New(&conf.Db)
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
