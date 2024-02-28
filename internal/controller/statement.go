package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/helper"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/service"
)

func GetStatements(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	statements, err := service.GetStatements(id)

	if err != nil {
		helper.WriteErrorResponse(rw, http.StatusNotFound, err)
		return
	}

	helper.WriteResponse(rw, http.StatusOK, statements)
}

func CreateStatement(rw http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))

	statement := entity.Statement{AccountId: id}

	if err := json.NewDecoder(req.Body).Decode(&statement); err != nil {
		helper.WriteErrorResponse(rw, http.StatusBadRequest, err)
		return
	}

	if !statement.Valid() {
		helper.WriteErrorResponse(rw, http.StatusUnprocessableEntity, fmt.Errorf("invalid statement"))
		return
	}

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

		helper.WriteErrorResponse(rw, statusCode, err)
		return
	}

	helper.WriteResponse(rw, http.StatusOK, account)
}
