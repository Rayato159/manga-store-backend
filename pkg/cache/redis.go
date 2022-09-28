package cache

import (
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/pkg/utils"
)

type RedisContext string

const (
	RedisConnection RedisContext = "RedisConnection"
)

func NewRedisConnection(cfg *configs.Configs) *redis.Client {
	ctx := context.WithValue(context.TODO(), RedisConnection, "Ctx.NewRedisConnection")
	defer ctx.Done()

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
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("error, can't connect to redis 😥")
		return nil
	}
	log.Printf("ping -> %v redis client has been connected 📕", pong)
	return rdb
}
