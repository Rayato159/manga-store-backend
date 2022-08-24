package entities

import "context"

type MonitorsContext string

const (
	MonitorsCon MonitorsContext = "MonitorsController"
	MonitorsUse MonitorsContext = "MonitorsUsecase"
	MonitorsRep MonitorsContext = "MonitorsRepository"
)

type MonitorsUsecase interface {
	HealthCheck(ctx context.Context) Monitor
	VersionCheck(ctx context.Context) Monitor
}

type Monitor struct {
	Health  string
	Version string
}
