package statement

import (
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"
	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence"
)

func Save(stt entity.Statement) (err error) {
	db, err := persistence.GetConnection()

	if err != nil {
		return
	}

	_, err = db.Exec(`INSERT INTO "Statement" ("Value", "Type", "Description", "Date") VALUES ($1, $2, $3, NOW())`, stt.Value, stt.Type, stt.Description)

	return
}
