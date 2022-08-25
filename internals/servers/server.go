package servers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/pkg/utils"
)

type Server struct {
	Fiber *fiber.App
	Cfg   *configs.Configs
	Db    *sqlx.DB
	File  *os.File
}

func NewServer(cfg *configs.Configs, db *sqlx.DB, file *os.File) *Server {
	fiberConfigs := configs.NewFiberConfig(cfg)
	app := fiber.New(fiberConfigs)
	return &Server{
		Fiber: app,
		Cfg:   cfg,
		Db:    db,
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

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.Fiber.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
