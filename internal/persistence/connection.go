package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/helper"
	_ "github.com/lib/pq"
)

var mutex = sync.Mutex{}
var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	mutex.Lock()

	defer mutex.Unlock()

	if db != nil {
		return db
	}

	connectionString := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	newConnection, err := sql.Open("postgres", connectionString)
	helper.CheckFatalError(err)

	db = newConnection

	err = db.Ping()
	helper.CheckFatalError(err)

	return db
}
