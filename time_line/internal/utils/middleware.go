package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authheader := c.Get("authorization")
		if authheader == "" || !strings.HasPrefix(authheader, "Bearer") {
			return c.Status(401).JSON(fiber.Map{
				"error": "missing or invalid token",
			})
		}
		tokenStirng := strings.TrimPrefix(authheader, "Bearer")

		userID, err := ParseJWT(tokenStirng)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid or experid token",
			})
		}
		c.Locals("user_id", userID)

		return c.Next()
	}
}
