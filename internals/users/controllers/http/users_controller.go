package http

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
)

type usersCon struct {
	UsersUse entities.UsersUsecase
	Cfg      *configs.Configs
}

func NewUsersController(r fiber.Router, cfg *configs.Configs, usersUse entities.UsersUsecase) {
	controller := &usersCon{
		UsersUse: usersUse,
		Cfg:      cfg,
	}
	r.Post("/", controller.Register)
}

func (uc *usersCon) Register(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.UsersCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.UsersCon).(int64)))

	req := new(entities.UsersRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.StatusBadRequest,
			Message:    "error, can't parse a request body into the struct",
			Result: entities.Result{
				Data: nil,
			},
		})
	}

	if req.Password != req.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.StatusBadRequest,
			Message:    "error, confirm password is not match",
			Result: entities.Result{
				Data: nil,
			},
		})
	}

	switch req.Role {
	case entities.Admin:
		if req.AdminKey != uc.Cfg.App.AdminKey {
			return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
				Status:     fiber.ErrBadRequest.Message,
				StatusCode: fiber.StatusBadRequest,
				Message:    "error, admin key is invalid",
				Result: entities.Result{
					Data: nil,
				},
			})
		}
	case entities.User:
	default:
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.StatusBadRequest,
			Message:    "error, role is invalid",
			Result: entities.Result{
				Data: nil,
			},
		})
	}

	res, err := uc.UsersUse.Register(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Response{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
			Result: entities.Result{
				Data: nil,
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.Response{
		Status:     "Created",
		StatusCode: fiber.StatusCreated,
		Message:    "",
		Result: entities.Result{
			Data: res,
		},
	})
}
