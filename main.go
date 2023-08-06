package main

import (
	"context"
	"fmt"
	"url-shortener/db"
	"url-shortener/handler"
	"url-shortener/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"go.uber.org/fx"
)

func fiberInstance(lc fx.Lifecycle, linkHandlers *handler.LinkHandler) *fiber.App {

	app := fiber.New()
	app.Get("/*", linkHandlers.GetLink)
	app.Post("/link", linkHandlers.CreateLink)
	
	app.Use(favicon.New())


	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port 3030")
			go app.Listen(":3030")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}


func main() {
	fx.New(
		fx.Provide(
			db.CreateSQLLiteConnection,
			repository.NewLinkRepository,
			handler.NewLinkHandler),
		fx.Invoke(fiberInstance),
	).Run()
}