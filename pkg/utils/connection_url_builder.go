package utils

import (
	"errors"
	"fmt"

	"github.com/rayato159/manga-store/configs"
)

func ConnectionUrlBuilder(stuff string, cfg *configs.Configs) (string, error) {
	var url string

	switch stuff {
	case "fiber":
		url = fmt.Sprintf("%s:%s", cfg.Fiber.Host, cfg.Fiber.Port)
	case "postgresql":
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.PostgreSQL.Host,
			cfg.PostgreSQL.Port,
			cfg.PostgreSQL.Username,
			cfg.PostgreSQL.Password,
			cfg.PostgreSQL.Database,
			cfg.PostgreSQL.SSLMode,
		)
	case "redis":
		url = fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}
	return url, nil
}
