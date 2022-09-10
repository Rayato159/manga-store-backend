package servers

import (
	"log"

	"github.com/rayato159/manga-store/internals/entities"

	_monitorsHttp "github.com/rayato159/manga-store/internals/monitors/controllers/http"
	_monitorsUsecase "github.com/rayato159/manga-store/internals/monitors/usecases"

	_usersHttp "github.com/rayato159/manga-store/internals/users/controllers/http"
	_usersRepository "github.com/rayato159/manga-store/internals/users/repositories"
	_usersUsecase "github.com/rayato159/manga-store/internals/users/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *Server) MapHandlers() error {
	// For save a log to file
	s.Fiber.Use(logger.New(logger.Config{
		Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Thailand/Bangkok",
		Output:     s.File,
	}))

	// For console a log
	s.Fiber.Use(logger.New(logger.Config{
		Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Thailand/Bangkok",
	}))

	// Group a version
	v1 := s.Fiber.Group("/v1")

	//* Monitors group.
	monitorsUsecase := _monitorsUsecase.NewMonitorsUsecase()
	_monitorsHttp.NewMonitorsController(v1, s.Cfg, monitorsUsecase)

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, s.Cfg, usersUsecase)

	// End point not found response
	s.Fiber.Use(func(c *fiber.Ctx) error {
		log.Println("error, endpoint is not found")
		return c.Status(fiber.StatusNotFound).JSON(entities.Response{
			Status:     fiber.ErrNotFound.Message,
			StatusCode: fiber.StatusNotFound,
			Message:    "error, endpoint is not found",
		})
	})

	return nil
}
