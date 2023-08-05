package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func DB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "file::memory:?cache=shared")
	if (err != nil) {
		log.Fatal(err)
	}

	return db, nil
}