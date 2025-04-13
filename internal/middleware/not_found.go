package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON("Route " + c.Path() + " not found")
	}
}
