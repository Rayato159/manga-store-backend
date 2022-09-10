package http

import (
	"context"
	"log"

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

func (tuc *usersCon) Register(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.UsersCon, "Con.TestRegister")
	defer log.Println(ctx.Value(entities.UsersCon))

	return nil
}
