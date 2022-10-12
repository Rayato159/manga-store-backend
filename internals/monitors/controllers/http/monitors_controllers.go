package http

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
)

type monitorsCon struct {
	MonitorsUse entities.MonitorsUsecase
	Cfg         *configs.Configs
}

func NewMonitorsController(r fiber.Router, cfg *configs.Configs, monitorsUse entities.MonitorsUsecase) {
	controller := &monitorsCon{
		MonitorsUse: monitorsUse,
		Cfg:         cfg,
	}
	r.Get("/", controller.HealthCheck)
}

func (mc *monitorsCon) HealthCheck(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.MonitorsCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.MonitorsCon).(int64)))

	res := mc.MonitorsUse.HealthCheck(ctx, mc.Cfg)
	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: res,
		},
	})
}
