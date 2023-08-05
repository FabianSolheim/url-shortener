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
	
	fmt.Println("Starting the server on :3000")
	log.Fatal(app.Listen(":3000"))
}