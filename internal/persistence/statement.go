package persistence

import "github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"

type StatementDAO entity.Statement

func (stt *StatementDAO) Save() (err error) {
	db, err := GetConnection()

	if err != nil {
		return
	}

	_, err = db.Exec(
		`INSERT INTO "Statement" ("Value", "Type", "Description", "Date") VALUES ($1, $2, $3, NOW())`,
		stt.Value,
		stt.Type,
		stt.Description,
	)

	return
}
