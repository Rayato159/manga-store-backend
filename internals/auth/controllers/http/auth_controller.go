package http

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/middlewares"
)

type authCon struct {
	AuthUse  entities.AuthUsecase
	UsersUse entities.UsersUsecase
	Cfg      *configs.Configs
	Redis    *redis.Client
}

func NewAuthController(r fiber.Router, cfg *configs.Configs, rdb *redis.Client, authUse entities.AuthUsecase, usersUse entities.UsersUsecase) {
	controller := &authCon{
		AuthUse:  authUse,
		UsersUse: usersUse,
		Cfg:      cfg,
		Redis:    rdb,
	}
	r.Post("/login", controller.Login)
	r.Post("/refresh-token", middlewares.JwtAuthentication(cfg), controller.RefreshToken)
}

func (ac *authCon) Login(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.AuthCon, "Con.Login")
	defer log.Println(ctx.Value(entities.AuthCon))

	req := new(entities.UsersCredentialsReq)
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

	res, err := ac.AuthUse.Login(ctx, ac.Cfg, ac.Redis, req)
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

	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: res,
		},
	})
}

func (ac *authCon) RefreshToken(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.AuthCon, "Con.RefreshToken")
	defer log.Println(ctx.Value(entities.AuthCon))

	req := new(entities.RefreshTokenReq)
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
	if req.RefreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.StatusBadRequest,
			Message:    "error, refresh token is invalid",
			Result: entities.Result{
				Data: nil,
			},
		})
	}

	res, err := ac.AuthUse.RefreshToken(ctx, ac.Cfg, ac.Redis, req.RefreshToken)
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

	return c.Status(fiber.StatusOK).JSON(entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: entities.Result{
			Data: res,
		},
	})
}
