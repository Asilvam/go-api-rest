package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	loadENV()
	user := os.Getenv("DBUSER")
	connStr := "postgres://" + user + ":hirefront@localhost:5432/hirefront?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}
