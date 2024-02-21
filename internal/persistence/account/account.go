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
