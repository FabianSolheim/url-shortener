package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"url-shortener/db"
	"url-shortener/handlers"
	"url-shortener/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func fiberInstance(lc fx.Lifecycle, linkHandlers *handlers.LinkHandler) *fiber.App {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()
	app.Get("/*", linkHandlers.GetLink)
	app.Post("/link", linkHandlers.CreateLink)


	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port:" + port)
			go app.Listen("0.0.0.0:" + port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}


func main() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file")
	}

	fx.New(
		fx.Provide(
			db.CreatePGConnection,
			repository.NewLinkRepository,
			handlers.NewLinkHandler),
		fx.Invoke(fiberInstance),
	).Run()
}