package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSQLLiteConnection() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", "file::memory:?cache=shared")
	err := db.Ping()
	if (err != nil) {
		log.Fatal(err)
	}
	db.Exec(`CREATE TABLE link (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		alias TEXT NOT NULL UNIQUE,
		link TEXT NOT NULL
	);`)

	return db
}