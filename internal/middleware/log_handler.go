package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

func LoggerInterceptor() fiber.Handler {
	return func(c *fiber.Ctx) error {

		log := logger.GetLogger()

		start := time.Now()

		err := c.Next()

		log.Info("HTTP Request",
			zap.String("time", time.Now().Format("15:04:05")),
			zap.Int("status", c.Response().StatusCode()),
			zap.String("latency", time.Since(start).String()),
			zap.String("ip", c.IP()),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		return err
	}
}
