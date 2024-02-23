package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/helper"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/service"
	_ "github.com/lib/pq"
)

func getStatements(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	statement := persistence.StatementDAO{}

	statements, _ := statement.GetLast10ByAccountId(id)

	res, _ := json.Marshal(statements)

	rw.Header().Add("content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(res))
}

func addStatement(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	statement := entity.Statement{AccountId: id}

	json.NewDecoder(req.Body).Decode(&statement)

	account, err := service.CreateStatement(statement)

	if err != nil {
		switch err {
		default:
			rw.WriteHeader(http.StatusInternalServerError)
		case helper.AccountNotFoundError:
			rw.WriteHeader(http.StatusNotFound)
		case helper.InsufficientBalanceError:
			rw.WriteHeader(http.StatusUnprocessableEntity)
		}

		fmt.Fprintf(rw, "{\"error\":\"%v\"}", err)
		return
	}

	res, _ := json.Marshal(account)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(res))
}

func main() {
	err := persistence.Connect()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /clientes/{id}/extrato", getStatements)
	http.HandleFunc("POST /clientes/{id}/transacoes", addStatement)

	addr := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	http.ListenAndServe(addr, nil)
}
