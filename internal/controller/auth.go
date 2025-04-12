package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/internal/dto"
	"github.com/vanthang24803/mini/internal/service"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

type AuthController struct {
	authService *service.AuthService
	log         *zap.Logger
	validate    *validator.Validate
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: service.NewAuthService(),
		log:         logger.GetLogger(),
		validate:    validator.New(),
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	c.log.Info("Login Handler")
	req := new(dto.LoginRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := c.validate.Struct(req); err != nil {
		c.log.Error("Validation failed", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	rs, err := c.authService.Login(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(rs)
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	req := new(dto.RegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := c.validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	rs, err := c.authService.Register(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.JSON(rs)
}
