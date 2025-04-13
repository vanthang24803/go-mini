package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/internal/service"
	"github.com/vanthang24803/mini/pkg/common"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

type MeController struct {
	meService *service.MeService
	log       *zap.Logger
	validate  *validator.Validate
}

func NewMeController() *MeController {
	return &MeController{
		meService: service.NewMeService(),
		log:       logger.GetLogger(),
		validate:  validator.New(),
	}
}

func (c *MeController) Profile(ctx *fiber.Ctx) error {

	payload, ok := ctx.Locals("info").(*common.JWTClaim)

	if !ok {
		c.log.Error("Unauthorized access attempt during logout")
		return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	rs, err := c.meService.Profile(payload.UserID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.JSON(rs)
}
