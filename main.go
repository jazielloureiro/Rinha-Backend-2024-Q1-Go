package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
	_ "github.com/lib/pq"
)

type Account struct {
	Id    int
	Limit int
	Value int
}

func handleStatements(rw http.ResponseWriter, req *http.Request) {
	row := persistence.GetConnection().QueryRow(`SELECT * FROM "Account" WHERE "Id" = $1`, req.PathValue("id"))

	err := row.Err()

	if err != nil {
		log.Fatal(err)
	}

	var account Account
	row.Scan(&account.Id, &account.Limit, &account.Value)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hello, %v!\n", account)
}

func main() {
	persistence.GetConnection()
	http.HandleFunc("GET /clientes/{id}/extrato", handleStatements)

	addr := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	http.ListenAndServe(addr, nil)
}
