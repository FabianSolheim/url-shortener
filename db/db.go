package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSQLLiteConnection() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", "_data.db")
	err := db.Ping()
	if (err != nil) {
		log.Fatal(err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS link (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		alias TEXT NOT NULL UNIQUE,
		link TEXT NOT NULL
	);`) //since i only have one table, i'll just create it here

	return db
}