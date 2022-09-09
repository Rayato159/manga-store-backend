package usecases

import (
	"context"
	"log"
	"time"

	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
)

type monitorsUse struct {
	MonitorsRepo any
}

func NewMonitorsUsecase() entities.MonitorsUsecase {
	return &monitorsUse{
		MonitorsRepo: nil,
	}
}

func (mu *monitorsUse) HealthCheck(ctx context.Context, cfg *configs.Configs) entities.Monitor {
	ctx = context.WithValue(ctx, entities.MonitorsUse, "Use.HealthCheck")
	defer log.Println(ctx.Value(entities.MonitorsUse))

	return entities.Monitor{
		Health:  "health is 100% 👌" + time.Now().Format("2006-01-02 15:04:05"),
		Version: cfg.App.Version,
	}
}
