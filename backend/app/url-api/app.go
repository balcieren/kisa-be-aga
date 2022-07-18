package urlapi

import (
	"backend/app/url-api/helper"
	v1 "backend/app/url-api/v1"
	"backend/pkg/config"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "Kisa-Be-Aga",
		ErrorHandler: helper.ErrorHandler,
	})
	return app
}

func Start(lifecycle fx.Lifecycle, app *fiber.App, c *config.Config, rv1 *v1.Router) {
	app.Use(logger.New())
	app.Use(cors.New())

	rv1.Setup()

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(fmt.Sprintf(":%s", c.UrlApiPort))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
