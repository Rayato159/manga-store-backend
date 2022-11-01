package servers

import (
	"log"

	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/middlewares"

	_monitorsHttp "github.com/rayato159/manga-store/internals/monitors/controllers/http"
	_monitorsUsecase "github.com/rayato159/manga-store/internals/monitors/usecases"

	_usersHttp "github.com/rayato159/manga-store/internals/users/controllers/http"
	_usersRepository "github.com/rayato159/manga-store/internals/users/repositories"
	_usersUsecase "github.com/rayato159/manga-store/internals/users/usecases"

	_authHttp "github.com/rayato159/manga-store/internals/auth/controllers/http"
	_authRepository "github.com/rayato159/manga-store/internals/auth/repositories"
	_authUsecase "github.com/rayato159/manga-store/internals/auth/usecases"

	_testsHttp "github.com/rayato159/manga-store/internals/tests/controllers/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	middlewares.NewCorsFiberHandler(s.Fiber)
	middlewares.NewFiberLoggerHandler(s.Fiber, s.File)

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

	//* Auth group
	authGroup := v1.Group("/auth")
	authRepository := _authRepository.NewAuthRepository(s.Db)
	authUsecase := _authUsecase.NewAuthUsecase(authRepository, usersRepository)
	_authHttp.NewAuthController(authGroup, s.Cfg, s.Redis, authUsecase, usersUsecase)

	//* Test group
	testsGroup := v1.Group("/tests")
	_testsHttp.NewTestsController(testsGroup, s.Cfg)

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
