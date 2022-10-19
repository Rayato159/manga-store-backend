package middlewares

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
)

func Authorization(roles ...entities.UsersRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.Context(), entities.MiddlewaresCon, time.Now().UnixMilli())
		log.Printf("called:\t%v", utils.Trace())
		defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.MiddlewaresCon).(int64)))

		rolesMap := map[entities.UsersRole]int{
			"user":    1, // 2^0
			"manager": 2, // 2^1
			"admin":   4, // 2^2
		}
		// Find a summation of expected role
		var sum int
		for i := range roles {
			sum += rolesMap[roles[i]]
		}
		roleBinary := utils.BinaryConvertor(sum, 2)

		// Get access token from cache of Fiber context
		userRole := rolesMap[entities.UsersRole(c.Locals("role").(string))]
		userRoleBinary := utils.BinaryConvertor(userRole, 2)

		// Bitwise operator to compare a role which can access or not
		for i := 0; i < 2; i++ {
			if roleBinary[i]&userRoleBinary[i] == 1 {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusUnauthorized).JSON(entities.Response{
			Status:     fiber.ErrUnauthorized.Message,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "error, have no permission to access",
			Result: entities.Result{
				Data: nil,
			},
		})
	}
}
