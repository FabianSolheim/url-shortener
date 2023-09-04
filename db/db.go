package db

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func CreatePGConnection() *sqlx.DB {
	connection_string := os.Getenv("POSTGRES_URL")

	db := sqlx.MustConnect("pgx", connection_string)
	
	err := db.Ping()
	if (err != nil) {
		log.Fatal(err)
	}
	err = goose.SetDialect("postgres")
    if err != nil {
        log.Fatalf("Error setting dialect: %v", err)
    }

    err = goose.Run("up", db.DB, "db/migrations")
    if err != nil {
        log.Fatalf("Error running migrations: %v", err)
    }

	return db
}