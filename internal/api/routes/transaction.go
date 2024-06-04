package routes

import (
	"log"
	"net/http"
)

func HandleDoTransaction(w http.ResponseWriter, r *http.Request) {

}

func HandleNewTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding new authorization...")
}

func HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed transactions...")
	w.Write([]byte("Transactions page"))
}
