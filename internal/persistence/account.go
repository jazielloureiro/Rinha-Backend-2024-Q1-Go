package persistence

import "github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"

type PostgresAccountRepository struct{}

func (par PostgresAccountRepository) Get(id int) (acc entity.Account, err error) {
	db, err := GetConnection()

	if err != nil {
		return
	}

	acc.Id = id
	row := db.QueryRow(`SELECT "Limit", "Value" FROM "Account" WHERE "Id" = $1`, acc.Id)

	err = row.Scan(&acc.Limit, &acc.Value)

	return
}

func (par PostgresAccountRepository) Update(acc *entity.Account, statementValue int) (err error) {
	db, err := GetConnection()

	if err != nil {
		return
	}

	row := db.QueryRow(
		`UPDATE "Account" SET "Value" = "Value" + $1 WHERE "Id" = $2 RETURNING "Value"`,
		statementValue,
		acc.Id,
	)

	err = row.Scan(&acc.Value)

	return
}

func NewAccountRepository() PostgresAccountRepository {
	return PostgresAccountRepository{}
}
