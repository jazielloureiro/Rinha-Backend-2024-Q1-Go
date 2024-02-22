package persistence

import "github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"

type AccountDAO entity.Account

func (acc *AccountDAO) Get() (err error) {
	db, err := GetConnection()

	if err != nil {
		return
	}

	row := db.QueryRow(`SELECT "Limit", "Value" FROM "Account" WHERE "Id" = $1`, acc.Id)

	err = row.Err()

	if err != nil {
		return
	}

	row.Scan(&acc.Limit, &acc.Value)

	return
}

func (acc *AccountDAO) Update() (err error) {
	db, err := GetConnection()

	if err != nil {
		return
	}

	row := db.QueryRow(`UPDATE "Account" SET "Value" = $1 WHERE "Id" = $2 RETURNING "Value"`, acc.Value, acc.Id)

	err = row.Err()

	if err != nil {
		return
	}

	row.Scan(&acc.Value)

	return
}
