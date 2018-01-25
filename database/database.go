package database

import (
	"database/sql"
	"log"
)

var database *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "go-native-rest.db")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	database = db

	startMigrations()
}

// GetConnection - database connection instance
func GetConnection() *sql.DB {
	return database
}

func startMigrations() {
	var db = GetConnection()

	createTableUser := `CREATE TABLE IF NOT EXISTS user (
												id INTEGER PRIMARY KEY,
												name TEXT NOT NULL
											);`
	_, err := db.Exec(createTableUser)
	if err != nil {
		log.Fatalln(err)
	}
}
