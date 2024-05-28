package main

import (
	"context"
	db "desafio-pic-pay/internal/storage/sqlc"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	testConnPol, _ := pgxpool.New(context.Background(), "postgresql://postgres:picpay-backend@localhost:5432/picpay-backend?sslmode=disable")

	testQueries := db.New(testConnPol)
	user, err := testQueries.CreateNewShopKeeperUser(context.Background(), db.CreateNewShopKeeperUserParams{
		FullName: "Jaw",
		Cpf:      "000.000.000-00",
		Email:    "tuyweber@gmail.com",
		Cnpj:     "61.687.803/0001-50",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
	// Get the isntance for Store
	//router := api.GetRouter()
	//
	//server := &http.Server{
	//	Addr:              ":42069",
	//	Handler:           router,
	//	ReadHeaderTimeout: 10 * time.Second,
	//	WriteTimeout:      10 * time.Second,
	//}
	//
	//log.Println("Server listening on port=localhost" + server.Addr)
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Fatal("Could not start server: ", err)
	//}
}
