package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/pkg/common"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

func HanlderAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {

		header := c.Get("Authorization")

		if header == "" {
			return c.Status(fiber.StatusUnauthorized).JSON("Missing authorization header")
		}

		currentToken := strings.Split(header, " ")

		if len(currentToken) != 2 || strings.ToLower(currentToken[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON("Invalid token")
		}

		token := currentToken[1]

		claims, err := common.ValidateToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON("Invalid token")
		}

		logger.GetLogger().Info("Claims", zap.Any("claims", claims))

		c.Locals("info", claims)

		return c.Next()
	}
}
