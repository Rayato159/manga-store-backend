package configs

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Configs struct {
	Fiber      Fiber
	PostgreSQL PostgreSQL
	Redis      Redis
	File       File
	App        App
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

type Redis struct {
	Host     string
	Port     string
	Password string
	Database string
}

type File struct {
	LogPath string
}

type App struct {
	Version                string
	AdminKey               string
	ManagerKey             string
	Stage                  string
	JwtSecretKey           string
	JwtAccessTokenExpires  string
	JwtRefreshTokenExpires string
	JwtSessionTokenExpires string
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
