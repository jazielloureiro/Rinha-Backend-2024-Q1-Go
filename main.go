package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence/account"
	_ "github.com/lib/pq"
)

func handleStatements(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	account, _ := account.Get(id)

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

	addr := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	http.ListenAndServe(addr, nil)
}
