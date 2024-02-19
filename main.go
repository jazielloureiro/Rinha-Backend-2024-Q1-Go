package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Account struct {
	Id    int
	Limit int
	Value int
}

func handleStatements(rw http.ResponseWriter, req *http.Request) {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT * FROM \"Account\" WHERE \"Id\" = $1", req.PathValue("id"))

	err = row.Err()

	if err != nil {
		log.Fatal(err)
	}

	var account Account
	row.Scan(&account.Id, &account.Limit, &account.Value)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hello, %v!\n", account)
}

func main() {
	http.HandleFunc("GET /clientes/{id}/extrato", handleStatements)

	addr := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	http.ListenAndServe(addr, nil)
}
