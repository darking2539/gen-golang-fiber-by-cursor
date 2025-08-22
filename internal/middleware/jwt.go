package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TokenParser interface {
	ParseToken(tokenString string) (interface{}, error)
}

type ClaimsExtractor func(interface{}) (string, error)

func JWTAuth(parse func(string) (interface{}, error), extract ClaimsExtractor) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := parse(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}
		username, err := extract(claims)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid claims"})
		}
		c.Locals("username", username)
		return c.Next()
	}
}
