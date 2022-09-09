package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/internals/entities"
)

type usersCon struct {
	UsersUC entities.UsersUsecase
}

func NewUsersController(r fiber.Router, usersUC entities.UsersUsecase) {
	controller := &usersCon{
		UsersUC: usersUC,
	}
	r.Post("/", controller.Register)
}

func (uc *usersCon) Register(c *fiber.Ctx) error {
	return nil
}
