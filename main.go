package main

import (
	"desafio-pic-pay/internal/api"
	"log"
	"net/http"
	"time"
)

func main() {
	router := api.GetRouter()

	server := &http.Server{
		Addr:              ":42069",
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	log.Println("Server listening on port=localhost" + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
