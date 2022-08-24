package configs

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Configs struct {
	Fiber      Fiber
	PostgreSQL PostgreSQL
}

type Fiber struct {
	Host                 string
	Port                 string
	ServerRequestTimeout string
}

// Database
type PostgreSQL struct {
	Host                   string
	Port                   string
	Protocol               string
	Username               string
	Password               string
	Database               string
	SSLMode                string
	MaxConnections         string
	MaxIdleConnections     string
	MaxLifeTimeConnections string
}

func NewFiberConfig(c *Configs) fiber.Config {
	readTimeoutSecondCount, _ := strconv.Atoi(c.Fiber.ServerRequestTimeout)

	// Time limit -> 30 seconds or setting in .env
	// Body limit -> 10 MiB
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondCount),
		BodyLimit:   10 * 1024 * 1024,
	}
}
