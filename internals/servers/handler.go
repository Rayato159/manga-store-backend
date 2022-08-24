package servers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *Server) MapHandlers(a *fiber.App) error {
	// For save a log to file
	a.Use(logger.New(logger.Config{
		Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Thailand/Bangkok",
		Output:     s.File,
	}))

	// For console a log
	a.Use(logger.New(logger.Config{
		Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Thailand/Bangkok",
	}))

	a.Use(func(c *fiber.Ctx) error {
		log.Println("error, endpoint is not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      fiber.ErrNotFound.Message,
			"status_code": fiber.StatusNotFound,
			"message":     "error, endpoint is not found",
		})
	})
	return nil
}
