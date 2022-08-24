package http

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/internals/entities"
)

type monitorsCon struct {
	MonitorsUC entities.MonitorsUsecase
}

func NewMonitorsController(r fiber.Router, monitorUC entities.MonitorsUsecase) {
	controller := &monitorsCon{
		MonitorsUC: monitorUC,
	}
	r.Get("/health", controller.HealthCheck)
	r.Get("/version", controller.VersionCheck)
}

func (mc *monitorsCon) HealthCheck(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.MonitorsCon, "Con.HealthCheck")
	defer log.Println(ctx.Value(entities.MonitorsCon))

	res := mc.MonitorsUC.HealthCheck(ctx)
	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: res.Health,
		},
	})
}

func (mc *monitorsCon) VersionCheck(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.MonitorsCon, "Con.VersionCheck")
	defer log.Println(ctx.Value(entities.MonitorsCon))

	res := mc.MonitorsUC.VersionCheck(ctx)
	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: res.Version,
		},
	})
}
