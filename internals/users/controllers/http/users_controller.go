package http

import (
	"context"
	"fmt"
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
	r.Patch("/:user_id/change-password", controller.ChangePassword)
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
		if req.Key != uc.Cfg.App.AdminKey {
			return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
				Status:     fiber.ErrBadRequest.Message,
				StatusCode: fiber.StatusBadRequest,
				Message:    "error, admin key is invalid",
				Result: entities.Result{
					Data: nil,
				},
			})
		}
	case entities.Manager:
		if req.Key != uc.Cfg.App.ManagerKey {
			return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
				Status:     fiber.ErrBadRequest.Message,
				StatusCode: fiber.StatusBadRequest,
				Message:    "error, manager key is invalid",
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

func (uc *usersCon) ChangePassword(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.UsersCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.UsersCon).(int64)))

	userId := c.Params("user_id")
	req := new(entities.ChangePasswordReq)
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
	req.UserId = userId

	if err := uc.UsersUse.ChangePassword(ctx, req); err != nil {
		switch err.Error() {
		case "error, old password is invalid":
			return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
				Status:     fiber.ErrBadRequest.Message,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
				Result: entities.Result{
					Data: nil,
				},
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(entities.Response{
				Status:     fiber.ErrInternalServerError.Message,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
				Result: entities.Result{
					Data: nil,
				},
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: fmt.Sprintf("success, password was changed for user_id -> %s", userId),
		},
	})
}
