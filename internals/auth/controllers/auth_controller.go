package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
)

type authCon struct {
	AuthUse entities.AuthUsecase
	Cfg     *configs.Configs
}

func NewAuthController(r fiber.Router, cfg *configs.Configs, authUse entities.AuthUsecase) {
	controller := &authCon{
		AuthUse: authUse,
		Cfg:     cfg,
	}
	r.Post("/login", controller.Login)
}

func (ac *authCon) Login(c *fiber.Ctx) error {
	return nil
}
