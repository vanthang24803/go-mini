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
	userCollection *mongo.Collection
	log            *zap.Logger
}

func NewMeService() *MeService {
	return &MeService{
		userCollection: database.GetCollection(constant.COLLECTION_USER),
		log:            logger.GetLogger(),
	}
}

func (s *MeService) Profile(userID string) (*entity.User, *exception.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var currentUser entity.User

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		s.log.Error("invalid user ID format", zap.String("userId", userID))
		return nil, exception.ERROR_INVALID_USER_ID
	}

	filter := primitive.M{"_id": objectID}
	err = s.userCollection.FindOne(ctx, filter).Decode(&currentUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Error("user not found", zap.String("userId", userID))
			return nil, exception.ERROR_USER_NOT_FOUND
		}
		s.log.Error("error finding user", zap.Error(err))
		return nil, exception.ERROR_INTERNAL_SERVER
	}

	return &currentUser, nil
}

func (s *MeService) UpdateProfile(userID string, jsonData *dto.UpdateProfileRequest) (*entity.User, *exception.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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

	err = s.userCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Error("user not found", zap.String("userId", userID))
			return nil, exception.ERROR_USER_NOT_FOUND
		}
		s.log.Error("error updating user", zap.Error(err))
		return nil, exception.ERROR_INTERNAL_SERVER
	}

	return &updatedUser, nil
}
