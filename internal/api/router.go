package api

import (
	"context"
	"desafio-pic-pay/internal/api/errors"
	"desafio-pic-pay/internal/api/routes"
	db "desafio-pic-pay/internal/storage/sqlc"
	"log"
	"net/http"
)

// ServerHandler TODO: Create server struct that recieves each interface service of our API to start modeling the architecture, in case of needing help https://www.youtube.com/watch?v=367qYRy39zw
type ServerHandler struct{}

type APIHandler func(http.ResponseWriter, *http.Request, *db.Queries, context.Context) *errors.APIErrorWrapper

func makeAPIHandler(handler APIHandler, querier *db.Queries) http.HandlerFunc {
	requestContext := context.Background()

	return func(w http.ResponseWriter, req *http.Request) {
		if err := handler(w, req, querier, requestContext); err != nil {
			routes.WriteJSON(w, err.Code, err.WrappedError)
		}
	}
}

func GetRouter(querier *db.Queries) *http.ServeMux {
	router := http.NewServeMux()
	mountRoutes(router, querier)
	return router
}

func mountRoutes(router *http.ServeMux, querier *db.Queries) {
	router.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("New client accessed API base page...")
		writer.Write([]byte("Connected on test API page!"))
	})
	router.HandleFunc("POST /transaction", routes.HandleNewTransaction)
	router.HandleFunc("GET /transaction", routes.HandleGetTransactions)
	router.HandleFunc("POST /user", makeAPIHandler(routes.HandleCreateNewUser, querier))
}
