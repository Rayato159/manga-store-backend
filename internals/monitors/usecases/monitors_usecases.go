package usecases

import (
	"context"
	"log"
	"time"

	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
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
	ctx = context.WithValue(ctx, entities.MonitorsUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.MonitorsUse).(int64)))

	return entities.Monitor{
		Health:  "health is 100% 👌" + time.Now().Format("2006-01-02 15:04:05"),
		Version: cfg.App.Version,
	}
}
