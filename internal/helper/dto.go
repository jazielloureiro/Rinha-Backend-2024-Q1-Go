package helper

import (
	"time"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
)

type BalanceDTO struct {
	Date  time.Time `json:"data_extrato"`
	Limit int       `json:"limite"`
	Value int       `json:"total"`
}

type StatementsDTO struct {
	Balance    BalanceDTO         `json:"saldo"`
	Statements []entity.Statement `json:"ultimas_transacoes"`
}
