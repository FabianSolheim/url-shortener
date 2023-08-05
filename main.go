package main

import (
	"fmt"
	"log"
	"url-shortener/routes"

	"github.com/gofiber/fiber/v2"
)



func main() {
	app := fiber.New()
	app.Get("/*", routes.MapHandler)
	app.Post("/link", routes.LinkHandler)
	
	fmt.Println("Starting the server on :8080")
	log.Fatal(app.Listen(":8080"))
}