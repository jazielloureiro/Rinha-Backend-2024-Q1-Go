package main

import (
	"fmt"
	"net/http"
	"os"
)

func handleStatements(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hello, %v!\n", req.PathValue("id"))
}

func main() {
	http.HandleFunc("GET /clientes/{id}/extrato", handleStatements)

	addr := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	http.ListenAndServe(addr, nil)
}
