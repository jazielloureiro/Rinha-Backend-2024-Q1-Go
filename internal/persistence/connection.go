package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var mutex = sync.Mutex{}
var db *sql.DB

func GetConnection() (*sql.DB, error) {
	if db == nil {
		return db, nil
	}

	mutex.Lock()

	defer mutex.Unlock()

	if db != nil {
		return db, nil
	}

	err := Connect()

	return db, err
}

func Connect() error {
	connectionString := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	newConnection, err := sql.Open("postgres", connectionString)

	if err != nil {
		return err
	}

	db = newConnection

	err = db.Ping()

	if err != nil {
		return err
	}

	return nil
}
