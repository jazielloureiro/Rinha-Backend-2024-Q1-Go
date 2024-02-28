package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
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

	if newConnection, err := sql.Open("postgres", connectionString); err != nil {
		return err
	} else {
		db = newConnection
	}

	if err := db.Ping(); err != nil {
		return err
	}

	if maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS")); err != nil {
		return err
	} else {
		db.SetMaxOpenConns(maxOpenConns)
	}

	if maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS")); err != nil {
		return err
	} else {
		db.SetMaxOpenConns(maxIdleConns)
	}

	return nil
}
