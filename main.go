package main

import (
	"context"
	"desafio-pic-pay/internal/api"
	db "desafio-pic-pay/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"time"
)

const (
	dbSource = "postgresql://postgres:picpay-backend@localhost:5432/picpay-backend?sslmode=disable"
)

func main() {
	pool, err := pgxpool.New(context.Background(), dbSource)
	queries := db.New(pool)
	router := api.GetRouter(queries)

	server := &http.Server{
		Addr:              ":42069",
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	log.Println("Server listening on port=localhost" + server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
