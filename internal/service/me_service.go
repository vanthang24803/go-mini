package service

import (
	"context"
	"time"

	"github.com/vanthang24803/mini/internal/dto"
	"github.com/vanthang24803/mini/internal/entity"
	"github.com/vanthang24803/mini/pkg/constant"
	"github.com/vanthang24803/mini/pkg/database"
	"github.com/vanthang24803/mini/pkg/exception"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MeService struct {
	ctx            context.Context
	log            *zap.Logger
	userCollection *mongo.Collection
	redis          database.Redis
}

func NewMeService() *MeService {
	return &MeService{
		ctx:            context.Background(),
		log:            logger.GetLogger(),
		redis:          database.NewRedisService(),
		userCollection: database.GetCollection(constant.COLLECTION_USER),
	}
}

func (s *MeService) Profile(userID string) (*entity.User, *exception.Error) {
	var currentUser entity.User

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		s.log.Error("invalid user ID format", zap.String("userId", userID))
		return nil, exception.ERROR_INVALID_USER_ID
	}

	err = s.redis.Get(s.ctx, userID, &currentUser)

	if err == nil {
		return &currentUser, nil
	}

	filter := primitive.M{"_id": objectID}
	err = s.userCollection.FindOne(s.ctx, filter).Decode(&currentUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Error("user not found in database", zap.String("userId", userID))
			return nil, exception.ERROR_USER_NOT_FOUND
		}
		s.log.Error("error finding user in database", zap.Error(err))
		return nil, exception.ERROR_INTERNAL_SERVER
	}

	err = s.redis.Set(s.ctx, userID, currentUser, time.Minute*5)

	if err != nil {
		s.log.Error("failed to set user in redis", zap.Error(err))
		return nil, exception.ERROR_INTERNAL_SERVER
	}

	return &currentUser, nil
}

func (s *MeService) UpdateProfile(userID string, jsonData *dto.UpdateProfileRequest) (*entity.User, *exception.Error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		s.log.Error("invalid user ID format", zap.String("userId", userID))
		return nil, exception.ERROR_INVALID_USER_ID
	}

	filter := primitive.M{"_id": objectID}
	update := primitive.M{
		"$set": primitive.M{
			"firstName":     jsonData.FirstName,
			"lastName":      jsonData.LastName,
			"phone":         jsonData.Phone,
			"address":       jsonData.Address,
			"gender":        jsonData.Gender,
			"date_of_birth": jsonData.DateOfBirth,
			"updatedAt":     time.Now(),
		},
	}

	var updatedUser entity.User
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = s.userCollection.FindOneAndUpdate(s.ctx, filter, update, opts).Decode(&updatedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Error("user not found", zap.String("userId", userID))
			return nil, exception.ERROR_USER_NOT_FOUND
		}
		s.log.Error("error updating user", zap.Error(err))
		return nil, exception.ERROR_INTERNAL_SERVER
	}

	_ = s.redis.Set(s.ctx, userID, updatedUser, time.Minute*5)

	return &updatedUser, nil
}

func (s *MeService) ActiveAccount(userID string) (bool, *exception.Error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		s.log.Error("invalid user ID format", zap.String("userId", userID))
		return false, exception.ERROR_INVALID_USER_ID
	}

	filter := primitive.M{"_id": objectID}
	update := primitive.M{
		"$set": primitive.M{
			"active":    true,
			"updatedAt": time.Now(),
		},
	}

	var updatedUser entity.User
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = s.userCollection.FindOneAndUpdate(s.ctx, filter, update, opts).Decode(&updatedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Error("user not found", zap.String("userId", userID))
			return false, exception.ERROR_USER_NOT_FOUND
		}
		s.log.Error("error updating user", zap.Error(err))
		return false, exception.ERROR_INTERNAL_SERVER
	}

	_ = s.redis.Set(s.ctx, userID, updatedUser, time.Minute*5)

	return true, nil

}
