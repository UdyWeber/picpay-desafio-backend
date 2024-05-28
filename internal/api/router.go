package api

import (
	"desafio-pic-pay/internal/api/routes"
	"log"
	"net/http"
)

func GetRouter() *http.ServeMux {
	router := http.NewServeMux()
	mountRoutes(router)
	return router
}

func mountRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("New client accessed API base page...")
		writer.Write([]byte("Connected on test API page!"))
	})
	router.HandleFunc("POST /transaction", routes.HandleNewTransaction)
	router.HandleFunc("GET /transaction", routes.HandleGetTransactions)
}
