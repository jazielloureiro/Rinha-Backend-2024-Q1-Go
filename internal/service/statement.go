package service

import (
	"time"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/helper"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
)

func CreateStatement(statement entity.Statement) (entity.Account, error) {
	account := persistence.AccountDAO{Id: statement.AccountId}

	if err := account.Get(); err != nil {
		return entity.Account{}, helper.AccountNotFoundError
	}

	statementValue := statement.Value

	if statement.Type == "d" {
		statementValue = -statementValue
	}

	account.Value += statementValue

	if err := account.Update(); err != nil {
		return entity.Account{}, helper.InsufficientBalanceError
	}

	statementDAO := persistence.StatementDAO(statement)

	statementDAO.Save()

	return entity.Account(account), nil
}

func GetStatements(accountId int) (helper.StatementsDTO, error) {
	account := persistence.AccountDAO{Id: accountId}

	if err := account.Get(); err != nil {
		return helper.StatementsDTO{}, helper.AccountNotFoundError
	}

	statementDAO := persistence.StatementDAO{}

	statements, _ := statementDAO.GetLast10ByAccountId(accountId)

	return helper.StatementsDTO{
		Balance: helper.BalanceDTO{
			Date:  time.Now(),
			Limit: account.Limit,
			Value: account.Value,
		},
		Statements: statements,
	}, nil
}
