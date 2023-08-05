package main

import (
	"fmt"
	"log"
	"url-shortener/db"
	"url-shortener/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`CREATE TABLE link (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			alias TEXT NOT NULL UNIQUE,
			link TEXT NOT NULL
	);`)

	app := fiber.New()
	app.Get("/*", routes.MapHandler)
	app.Post("/link", routes.LinkHandler)

	fmt.Println("Starting the server on :3030")
	log.Fatal(app.Listen(":3030"))
}
