package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

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
		Device:    userAgent,
		RequestID: requestID,
	}

	return c.Status(code).JSON(&Response{
		Status:   code,
		Success:  false,
		Error:    err.Error(),
		Metadata: metadata,
	})
}
