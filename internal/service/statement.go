package service

import (
	"time"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/helper"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
)

type StatementService struct {
	accountRepository persistence.AccountRepository
}

func (ss StatementService) Save(statement entity.Statement) (entity.Account, error) {
	account, err := ss.accountRepository.Get(statement.AccountId)

	if err != nil {
		return account, helper.AccountNotFoundError
	}

	statementValue := statement.Value

	if statement.Type == "d" {
		statementValue = -statementValue
	}

	if err := ss.accountRepository.Update(&account, statementValue); err != nil {
		return account, helper.InsufficientBalanceError
	}

	statementDAO := persistence.StatementDAO(statement)

	statementDAO.Save()

	return entity.Account(account), nil
}

func (ss StatementService) Get(accountId int) (helper.StatementsDTO, error) {
	account, err := ss.accountRepository.Get(accountId)

	if err != nil {
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

func NewStatementService() StatementService {
	return StatementService{
		accountRepository: persistence.NewAccountRepository(),
	}
}
