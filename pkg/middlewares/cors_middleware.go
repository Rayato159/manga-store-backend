package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewCorsFiberHandler(app *fiber.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
			AllowCredentials: true,
		}))
		return c.Next()
	}
}
