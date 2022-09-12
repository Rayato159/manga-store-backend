package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
)

type authCon struct {
	AuthUse  entities.AuthUsecase
	UsersUse entities.UsersUsecase
	Cfg      *configs.Configs
}

func NewAuthController(r fiber.Router, cfg *configs.Configs, authUse entities.AuthUsecase, usersUse entities.UsersUsecase) {
	controller := &authCon{
		AuthUse:  authUse,
		UsersUse: usersUse,
		Cfg:      cfg,
	}
	r.Post("/login", controller.Login)
}

func (ac *authCon) Login(c *fiber.Ctx) error {
	return nil
}
