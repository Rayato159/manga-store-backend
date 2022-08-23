package configs

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Configs struct {
	Fiber    Fiber
	Database PostgreSQL
}

type Fiber struct {
	Host              string
	Port              string
	ServerReadTimeout string
}

// Database
type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
}

func NewFiberConfig(c *Configs) fiber.Config {
	readTimeoutSecondCount, _ := strconv.Atoi(c.Fiber.ServerReadTimeout)

	// Time limit -> 60 seconds or setting in .env
	// Body limit -> 10 MiB
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondCount),
		BodyLimit:   10 * 1024 * 1024,
	}
}
