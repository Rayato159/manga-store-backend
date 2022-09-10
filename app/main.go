package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/servers"
	"github.com/rayato159/manga-store/pkg/cache"
	"github.com/rayato159/manga-store/pkg/databases"
	"github.com/rayato159/manga-store/pkg/utils"
)

func main() {
	// Load dotenv config
	utils.LoadDotenv(os.Args[1])
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.Fiber.Host = os.Getenv("FIBER_HOST")
	cfg.Fiber.Port = os.Getenv("FIBER_PORT")
	cfg.Fiber.ServerRequestTimeout = os.Getenv("FIBER_REQUEST_TIMEOUT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// Redis
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Port = os.Getenv("REDIS_PORT")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
	cfg.Redis.Database = os.Getenv("REDIS_DATABASE")

	// App
	cfg.App.Version = os.Getenv("APP_VERSION")
	cfg.App.AdminKey = os.Getenv("ADMIN_KEY")
	cfg.App.Stage = os.Getenv("STAGE")

	// File
	cfg.File.LogPath = os.Getenv("FILE_LOG_PATH")

	// New Database
	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	// Log File
	filePath := fmt.Sprintf("%v/%v.log", cfg.File.LogPath, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error, opening file: %v", err)
	}
	defer file.Close()

	rdb := cache.NewRedisConnection(cfg)
	defer rdb.Conn().Close()

	s := servers.NewServer(cfg, db, rdb, file)
	s.Start()
}
