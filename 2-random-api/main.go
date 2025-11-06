package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {
	port := ":8081"
	router := http.NewServeMux()
	router.HandleFunc("/", handlerRandom)
	server := http.Server{
		Addr:    port,
		Handler: router,
	}
	server.ListenAndServe()
}

func handlerRandom(w http.ResponseWriter, r *http.Request) {
	num := rand.IntN(6) + 1
	w.Write([]byte(fmt.Sprint(num)))
}
