package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
	_ "github.com/lib/pq"
)

func handleStatements(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	account := persistence.AccountDAO{Id: id}

	account.Get()

	res, _ := json.Marshal(account)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(res))
}

func addStatement(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	statement := persistence.StatementDAO{AccountId: id}

	json.NewDecoder(req.Body).Decode(&statement)

	account := persistence.AccountDAO{Id: id}

	if err := account.Get(); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, err)
		return
	}

	account.Value += statement.Value

	if err := account.Update(); err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(rw, err)
		return
	}

	statement.Save()

	res, _ := json.Marshal(account)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(res))
}

func main() {
	err := persistence.Connect()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /clientes/{id}/extrato", handleStatements)
	http.HandleFunc("POST /clientes/{id}/transacoes", addStatement)

	addr := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	http.ListenAndServe(addr, nil)
}
