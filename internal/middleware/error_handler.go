package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/pkg/util"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&util.BaseResponse{
		Status:  fiber.StatusNotFound,
		Success: false,
		Error:   err.Error(),
		Metadata: util.Metadata{
			Timestamp: time.Now().UTC(),
			Version:   "v1.0",
		},
	})
}
