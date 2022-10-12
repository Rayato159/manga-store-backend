package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/middlewares"
)

type testsCon struct {
	Cfg *configs.Configs
}

func NewTestsController(r fiber.Router, cfg *configs.Configs) {
	controller := &testsCon{
		Cfg: cfg,
	}
	r.Get("/authentication", middlewares.JwtAuthentication(controller.Cfg), controller.Authentication)
	r.Get("/authorization", middlewares.JwtAuthentication(controller.Cfg), middlewares.Authorization("admin"), controller.Authorization)
}

func (tc *testsCon) Authentication(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: nil,
		},
	})
}

func (tc *testsCon) Authorization(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: nil,
		},
	})
}
