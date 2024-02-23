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

	statements := service.GetStatements(id)

	helper.WriteResponse(rw, http.StatusOK, statements)
}

func addStatement(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	statement := entity.Statement{AccountId: id}

	json.NewDecoder(req.Body).Decode(&statement)

	account, err := service.CreateStatement(statement)

	if err != nil {
		var statusCode int

		switch err {
		default:
			statusCode = http.StatusInternalServerError
		case helper.AccountNotFoundError:
			statusCode = http.StatusNotFound
		case helper.InsufficientBalanceError:
			statusCode = http.StatusUnprocessableEntity
		}

		helper.WriteResponse(
			rw,
			statusCode,
			map[string]string{"error": err.Error()},
		)

		return
	}

	helper.WriteResponse(rw, http.StatusOK, account)
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
