package main

import (
	"fmt"
	"net/http"
	"math/rand/v2"
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
	num := rand.IntN(6)
	w.Write([]byte(fmt.Sprint(num)))
}
