package service

import (
	"fmt"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
)

func CreateStatement(statement entity.Statement) (entity.Account, error) {
	account := persistence.AccountDAO{Id: statement.AccountId}

	if err := account.Get(); err != nil {
		return entity.Account{}, fmt.Errorf("there's not account for the given id")
	}

	statementValue := statement.Value

	if statement.Type == "d" {
		statementValue = -statementValue
	}

	account.Value += statementValue

	if err := account.Update(); err != nil {
		return entity.Account{}, fmt.Errorf("insufficient balance")
	}

	statementDAO := persistence.StatementDAO(statement)

	statementDAO.Save()

	return entity.Account(account), nil
}
