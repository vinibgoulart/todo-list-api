package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDatabase() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	envFile, _ := godotenv.Read(".env")

	dbUser := envFile["DB_USER"]
	dbPassword := envFile["DB_PASSWORD"]
	dbName := envFile["DB_NAME"]
	dbPort := envFile["DB_PORT"]
	dbHost := envFile["DB_HOST"]

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(15 * time.Second)
	db.SetConnMaxLifetime(15 * time.Second)

	return db, nil
}
