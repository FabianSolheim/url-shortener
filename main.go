package main

import (
	"context"
	"fmt"
	"url-shortener/db"
	"url-shortener/handlers"
	"url-shortener/repository"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func fiberInstance(lc fx.Lifecycle, linkHandlers *handlers.LinkHandler) *fiber.App {

	app := fiber.New()
	app.Get("/*", linkHandlers.GetLink)
	app.Post("/link", linkHandlers.CreateLink)


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
			handlers.NewLinkHandler),
		fx.Invoke(fiberInstance),
	).Run()
}