package middlewares

import (
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
)

func JwtAuthentication(cfg *configs.Configs) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.Context(), entities.AuthCon, "Middleware.JwtAuthentication")
		defer log.Println(ctx.Value(entities.AuthCon))

		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			log.Println("error, authorization header is empty.")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":      "Unauthorized",
				"status_code": fiber.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
		}

		userId, err := utils.JwtExtractPayload(ctx, cfg, "user_id", accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.Response{
				Status:     fiber.ErrUnauthorized.Message,
				StatusCode: fiber.StatusUnauthorized,
				Message:    err.Error(),
				Result: entities.Result{
					Data: nil,
				},
			})
		}

		paramsUserId := c.Params("user_id")
		if paramsUserId != "" && paramsUserId != userId {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.Response{
				Status:     fiber.ErrUnauthorized.Message,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "error, have no permission to access",
				Result: entities.Result{
					Data: nil,
				},
			})
		}
		return c.Next()
	}
}
