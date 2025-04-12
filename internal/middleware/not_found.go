package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NotFoundHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		requestID := c.Get("X-Request-ID")
		userAgent := c.Get("User-Agent")

		if requestID == "" {
			requestID = uuid.New().String()
		}

		metadata := Metadata{
			Timestamp: time.Now(),
			Version:   "1.0",
			Path:      c.Path(),
			Method:    c.Method(),
			RequestID: requestID,
			Device:    userAgent,
		}

		return c.Status(fiber.StatusNotFound).JSON(&Response{
			Status:   fiber.StatusNotFound,
			Success:  false,
			Error:    "Route " + c.Path() + " not found",
			Metadata: metadata,
		})
	}
}
