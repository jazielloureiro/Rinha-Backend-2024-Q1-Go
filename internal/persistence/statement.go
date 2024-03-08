package persistence

import "github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"

type PostgresStatementRepository struct{}

func (psr PostgresStatementRepository) Save(stt entity.Statement) (err error) {
	db, err := GetConnection()

	if err != nil {
		return
	}

	_, err = db.Exec(
		`INSERT INTO "Statement" ("AccountId", "Value", "Type", "Description", "Date") VALUES ($1, $2, $3, $4, NOW())`,
		stt.AccountId,
		stt.Value,
		stt.Type,
		stt.Description,
	)

	return
}

func (psr PostgresStatementRepository) GetLast10FromAccount(accountId int) (stts []entity.Statement, err error) {
	db, err := GetConnection()

	if err != nil {
		return
	}

	rows, err := db.Query(
		`SELECT "Value", "Type", "Description", "Date" FROM "Statement" WHERE "AccountId" = $1 ORDER BY "Id" DESC LIMIT 10`,
		accountId,
	)

	if err != nil {
		return
	}

	for i := 0; rows.Next(); i++ {
		var s entity.Statement

		rows.Scan(&s.Value, &s.Type, &s.Description, &s.Date)

		stts = append(stts, s)
	}

	return
}

func NewStatementRepository() PostgresStatementRepository {
	return PostgresStatementRepository{}
}
