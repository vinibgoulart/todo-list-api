package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/vinibgoulart/todo-list/lib"
)

var db *sql.DB

func ConnectDatabase() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName)

	logger := lib.NewLogger("main")

	logger.Info("Connecting to database...")
	logger.Info(dbInfo)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(15 * time.Second)
	db.SetConnMaxLifetime(15 * time.Second)

	return db, nil
}
