package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/pkg/util"
)

func NotFoundHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(util.NotFoundError("Route " + c.Path() + " not found"))
	}
}
