package service

import (
	"time"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/helper"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
)

type StatementService struct {
	accountRepository   persistence.AccountRepository
	statementRepository persistence.StatementRepository
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

	err = ss.statementRepository.Save(statement)

	return entity.Account(account), err
}

func (ss StatementService) Get(accountId int) (helper.StatementsDTO, error) {
	account, err := ss.accountRepository.Get(accountId)

	if err != nil {
		return helper.StatementsDTO{}, helper.AccountNotFoundError
	}

	statements, err := ss.statementRepository.GetLast10FromAccount(accountId)

	return helper.StatementsDTO{
		Balance: helper.BalanceDTO{
			Date:  time.Now(),
			Limit: account.Limit,
			Value: account.Value,
		},
		Statements: statements,
	}, err
}

func NewStatementService() StatementService {
	return StatementService{
		accountRepository:   persistence.NewAccountRepository(),
		statementRepository: persistence.NewStatementRepository(),
	}
}
