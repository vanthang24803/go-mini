package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/internal/dto"
	"github.com/vanthang24803/mini/internal/service"
	"github.com/vanthang24803/mini/pkg/common"
	"github.com/vanthang24803/mini/pkg/exception"
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
		c.log.Error("unauthorized access attempt during logout")
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.ERROR_CODE_UNAUTHORIZED)
	}

	rs, err := c.meService.Profile(payload.UserID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.JSON(rs)
}

func (c *MeController) UpdateProfile(ctx *fiber.Ctx) error {

	payload, ok := ctx.Locals("info").(*common.JWTClaim)

	if !ok {
		c.log.Error("unauthorized access attempt during logout")
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.ERROR_CODE_UNAUTHORIZED)
	}

	jsonData := new(dto.UpdateProfileRequest)

	if err := ctx.BodyParser(jsonData); err != nil {
		c.log.Error("error parsing request body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := c.validate.Struct(jsonData); err != nil {
		c.log.Error("validation error", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err := c.meService.UpdateProfile(payload.UserID, jsonData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.JSON("Update profile successfully!")
}
