package servers

import (
	"log"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/pkg/utils"
)

type Server struct {
	Fiber *fiber.App
	Cfg   *configs.Configs
	Db    *sqlx.DB
	Redis *redis.Client
	File  *os.File
}

func NewServer(cfg *configs.Configs, db *sqlx.DB, rdb *redis.Client, file *os.File) *Server {
	fiberConfigs := configs.NewFiberConfig(cfg)
	app := fiber.New(fiberConfigs)
	return &Server{
		Fiber: app,
		Cfg:   cfg,
		Db:    db,
		Redis: rdb,
		File:  file,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.Fiber.Host
	port := s.Cfg.Fiber.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.Fiber.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
