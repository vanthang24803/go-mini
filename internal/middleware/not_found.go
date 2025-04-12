package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/pkg/util"
)

func NotFoundHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&util.BaseResponse{
			Status:  fiber.StatusNotFound,
			Success: false,
			Error:   "Route " + c.Path() + " not found",
			Metadata: util.Metadata{
				Timestamp: time.Now().UTC(),
				Version:   "v1.0",
			},
		})
	}
}
