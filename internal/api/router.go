package api

import (
	"context"
	"desafio-pic-pay/internal/api/routes"
	db "desafio-pic-pay/internal/storage/sqlc"
	"log"
	"net/http"
)

// TODO: Create function that handles APIHandler erros and make them as a normal http.HandleFunc
type APIHandler func(http.ResponseWriter, *http.Request, db.Querier, context.Context) error

func GetRouter(querier db.Querier) *http.ServeMux {
	router := http.NewServeMux()
	mountRoutes(router, querier)
	return router
}

func mountRoutes(router *http.ServeMux, querier db.Querier) {
	router.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("New client accessed API base page...")
		writer.Write([]byte("Connected on test API page!"))
	})
	router.HandleFunc("POST /transaction", routes.HandleNewTransaction)
	router.HandleFunc("GET /transaction", routes.HandleGetTransactions)
	router.HandleFunc("POST /user", routes.HandleCreateNewUser(querier))
}
