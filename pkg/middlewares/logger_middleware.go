package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberLoggerHandler(app *fiber.App, fs *os.File) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// For save a log to file
		app.Use(logger.New(logger.Config{
			Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
			TimeFormat: "2006-01-02T15:04:05",
			TimeZone:   "Thailand/Bangkok",
			Output:     fs,
		}))

		// For console a log
		app.Use(logger.New(logger.Config{
			Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
			TimeFormat: "2006-01-02T15:04:05",
			TimeZone:   "Thailand/Bangkok",
		}))

		return c.Next()
	}
}
