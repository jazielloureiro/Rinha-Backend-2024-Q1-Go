package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/controller"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
	_ "github.com/lib/pq"
)

func main() {
	if err := persistence.Connect(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /clientes/{id}/extrato", controller.GetStatements)
	http.HandleFunc("POST /clientes/{id}/transacoes", controller.CreateStatement)

	addr := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	log.Fatal(http.ListenAndServe(addr, nil))
}
