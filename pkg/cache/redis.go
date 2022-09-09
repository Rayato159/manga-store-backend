package cache

import (
	"log"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/pkg/utils"
)

func NewRedisConnection(cfg *configs.Configs) *redis.Client {
	url, err := utils.ConnectionUrlBuilder("redis", cfg)
	if err != nil {
		panic(err.Error())
	}

	db, err := strconv.Atoi(cfg.Redis.Database)
	if err != nil {
		panic(err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: cfg.Redis.Password,
		DB:       db,
	})
	log.Println("redis client has been connected 📕")
	return rdb
}
