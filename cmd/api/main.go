package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vanthang24803/mini/internal/config"
	"github.com/vanthang24803/mini/internal/middleware"
	"github.com/vanthang24803/mini/internal/router"
	"github.com/vanthang24803/mini/pkg/database"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()

	logger.Init()
	log := logger.GetLogger()

	defer log.Sync()
	err := database.InitMongoDB(cfg)

	if err != nil {
		log.Fatal("failed to connect to MongoDB", zap.Error(err))
	}

	app := fiber.New(fiber.Config{
		AppName:               cfg.AppName,
		DisableStartupMessage: true,
		ErrorHandler:          middleware.ErrorHandler,
	})

	app.Use(fiberLogger.New())
	app.Use(middleware.SuccessHandler)
	app.Use(recover.New())

	router.SetupRoutes(app)

	app.Use(middleware.NotFoundHandler())

	log.Info(fmt.Sprintf("Server %s started on port %s 🚀", cfg.AppName, cfg.Server.Port))

	if err := app.Listen(cfg.Server.Port); err != nil {
		log.Fatal("server failed to start", zap.Error(err))
	}
}
