package entities

import (
	"context"

	"github.com/rayato159/manga-store/configs"
)

type MonitorsContext string

const (
	MonitorsCon MonitorsContext = "MonitorsController"
	MonitorsUse MonitorsContext = "MonitorsUsecase"
	MonitorsRep MonitorsContext = "MonitorsRepository"
)

type MonitorsUsecase interface {
	HealthCheck(ctx context.Context, cfg *configs.Configs) Monitor
}

type Monitor struct {
	Health  string
	Version string
}
