package account

import (
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
)

func Get(id int) (acc entity.Account, err error) {
	db, err := persistence.GetConnection()

	if err != nil {
		return
	}

	row := db.QueryRow(`SELECT "Id", "Limit", "Value" FROM "Account" WHERE "Id" = $1`, id)

	err = row.Err()

	if err != nil {
		return
	}

	row.Scan(&acc.Id, &acc.Limit, &acc.Value)

	return
}

func Update(acc entity.Account) (newAcc entity.Account, err error) {
	newAcc = acc
	db, err := persistence.GetConnection()

	if err != nil {
		return
	}

	row := db.QueryRow(`UPDATE "Account" SET "Value" = $1 WHERE "Id" = $2 RETURNING "Value"`, acc.Value, acc.Id)

	err = row.Err()

	if err != nil {
		return
	}

	row.Scan(&newAcc.Value)

	return
}
