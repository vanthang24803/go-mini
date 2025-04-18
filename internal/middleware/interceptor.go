package middleware

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Response struct {
	Status   int      `json:"status"`
	Success  bool     `json:"success"`
	Data     any      `json:"data,omitempty"`
	Error    any      `json:"error,omitempty"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Version   string    `json:"version"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	RequestID string    `json:"request_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Device    string    `json:"device,omitempty"`
}

func InterceptorHandler(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}

	status := c.Response().StatusCode()

	if status == fiber.StatusOK {
		var originalData any
		if err := json.Unmarshal(c.Response().Body(), &originalData); err != nil {
			return nil
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

		return c.JSON(&Response{
			Status:   fiber.StatusOK,
			Success:  true,
			Data:     originalData,
			Metadata: metadata,
		})
	} else if status >= fiber.StatusBadRequest {
		var originalData any
		if err := json.Unmarshal(c.Response().Body(), &originalData); err != nil {
			return nil
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

		return c.JSON(&Response{
			Status:   status,
			Success:  false,
			Error:    originalData,
			Metadata: metadata,
		})
	}

	return nil
}
