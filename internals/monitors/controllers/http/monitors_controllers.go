package http

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
)

type monitorsCon struct {
	MonitorsUC entities.MonitorsUsecase
	Cfg        *configs.Configs
}

func NewMonitorsController(r fiber.Router, cfg *configs.Configs, monitorUC entities.MonitorsUsecase) {
	controller := &monitorsCon{
		MonitorsUC: monitorUC,
		Cfg:        cfg,
	}
	r.Get("/", controller.HealthCheck)
}

func (mc *monitorsCon) HealthCheck(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.MonitorsCon, "Con.HealthCheck")
	defer log.Println(ctx.Value(entities.MonitorsCon))

	res := mc.MonitorsUC.HealthCheck(ctx, mc.Cfg)
	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: res,
		},
	})
}
