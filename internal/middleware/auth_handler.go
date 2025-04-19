package middleware

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	lo "github.com/samber/lo"
	"github.com/vanthang24803/mini/internal/entity"
	"github.com/vanthang24803/mini/pkg/common"
	"github.com/vanthang24803/mini/pkg/constant"
	"github.com/vanthang24803/mini/pkg/database"
	"github.com/vanthang24803/mini/pkg/exception"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func HandlerAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {

		header := c.Get("Authorization")

		if header == "" {
			return c.Status(fiber.StatusUnauthorized).JSON("Missing authorization header")
		}

		currentToken := strings.Split(header, " ")

		if len(currentToken) != 2 || strings.ToLower(currentToken[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON("Invalid token")
		}

		token := currentToken[1]

		claims, err := common.ValidateToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON("Invalid token")
		}

		logger.GetLogger().Info("Claims", zap.Any("claims", claims))

		c.Locals("info", claims)

		return c.Next()
	}
}

func HandlerAuthorization(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := database.GetCollection(constant.COLLECTION_USER)

		payload, ok := c.Locals("info").(*common.JWTClaim)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(exception.ERROR_CODE_UNAUTHORIZED)
		}

		objectID, err := primitive.ObjectIDFromHex(payload.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(exception.ERROR_CODE_UNAUTHORIZED)
		}

		var user entity.User
		err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(exception.ERROR_CODE_UNAUTHORIZED)
		}

		hasRole := lo.SomeBy(user.Roles, func(role string) bool {
			return lo.Contains(allowedRoles, role)
		})
		if !hasRole {
			return c.Status(fiber.StatusForbidden).JSON(exception.ERROR_CODE_FORBIDDEN)
		}

		return c.Next()
	}
}
