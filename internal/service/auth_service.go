package service

import (
	"context"
	"time"

	"github.com/vanthang24803/mini/internal/dto"
	"github.com/vanthang24803/mini/internal/entity"
	"github.com/vanthang24803/mini/pkg/common"
	"github.com/vanthang24803/mini/pkg/constant"
	"github.com/vanthang24803/mini/pkg/database"
	"github.com/vanthang24803/mini/pkg/exception"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type AuthService struct {
	userCollection  *mongo.Collection
	tokenCollection *mongo.Collection
	log             *zap.Logger
}

func NewAuthService() *AuthService {
	return &AuthService{
		userCollection:  database.GetCollection(constant.COLLECTION_USER),
		tokenCollection: database.GetCollection(constant.COLLECTION_TOKEN),
		log:             logger.GetLogger(),
	}
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.TokenResponse, *exception.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entity.User
	err := s.userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		s.log.Error("failed to find user by email", zap.Error(err))
		if err == mongo.ErrNoDocuments {
			return nil, exception.ERROR_INVALID_CREDENTIAL
		}
		return nil, exception.ERROR_INVALID_CREDENTIAL
	}

	if errComparePassword := common.ComparePassword(req.Password, user.HashedPassword); errComparePassword != nil {
		s.log.Error("invalid password", zap.Error(errComparePassword))
		return nil, exception.ERROR_INVALID_CREDENTIAL
	}

	token, refreshToken, err := common.GenerateJWT(user.ID.Hex(), user.Username)
	if err != nil {
		s.log.Error("failed to generate JWT", zap.Error(err))
		return nil, exception.ERROR_GENERATE_TOKEN
	}

	var currentToken entity.Token
	err = s.tokenCollection.FindOne(ctx, bson.M{
		"user_id": user.ID,
		"type":    constant.REFRESH_TOKEN,
	}, options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})).Decode(&currentToken)

	exp := time.Now().Add(7 * 24 * time.Hour).Unix()
	response := &dto.TokenResponse{
		Type:  "Bearer",
		Exp:   exp,
		Token: token,
	}

	if err == mongo.ErrNoDocuments {
		newToken := entity.Token{
			UserID:    user.ID,
			Token:     refreshToken,
			Type:      constant.REFRESH_TOKEN,
			ExpiresAt: time.Now().Add(constant.REFRESH_TOKEN_EXPIRATION),
			CreatedAt: time.Now(),
			Checksum:  common.GenerateCheckSum(&user),
		}

		if _, err := s.tokenCollection.InsertOne(ctx, newToken); err != nil {
			s.log.Error("failed to insert token", zap.Error(err))
			return nil, exception.ERROR_INSERT_TOKEN
		}

		response.RefreshToken = refreshToken
		return response, nil
	}

	if !common.CompareWithNow(currentToken.ExpiresAt) {
		update := bson.M{
			"$set": bson.M{
				"token":      refreshToken,
				"expires_at": time.Now().Add(30 * 24 * time.Hour),
			},
		}

		_, err := s.tokenCollection.UpdateOne(ctx,
			bson.M{"_id": currentToken.ID},
			update,
		)
		if err != nil {
			s.log.Error("failed to update token", zap.Error(err))
			return nil, exception.ERROR_INSERT_TOKEN
		}

		response.RefreshToken = refreshToken
		return response, nil
	}

	response.RefreshToken = currentToken.Token
	return response, nil
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*string, *exception.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser := &entity.User{}
	err := s.userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(existingUser)
	if err == nil {
		s.log.Error("email already exists", zap.String("email", req.Email))
		return nil, exception.ERROR_EMAIL_EXISTED
	}

	err = s.userCollection.FindOne(ctx, bson.M{"username": req.Username}).Decode(existingUser)
	if err == nil {
		s.log.Error("username already exists", zap.String("username", req.Username))
		return nil, exception.ERROR_CODE_USERNAME_EXISTED
	}
	if err != mongo.ErrNoDocuments {
		s.log.Error("error checking existing user", zap.Error(err))
		return nil, exception.ERROR_NO_DOCUMENT
	}

	hash, errHash := common.HashPassword(req.Password)
	if errHash != nil {
		s.log.Error("failed to hash password", zap.Error(errHash))
		return nil, exception.ERROR_HASH_PASSWORD
	}

	user := entity.User{
		Email:          req.Email,
		Username:       req.Username,
		HashedPassword: hash,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Role:           constant.ROLE_ROOT,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Timezone:       "Asia/Ho_Chi_Minh",
		Version:        constant.VERSION,
	}

	result, err := s.userCollection.InsertOne(ctx, user)
	if err != nil {
		s.log.Error("failed to insert user", zap.Error(err))
		return nil, &exception.Error{
			Code:    0,
			Message: err.Error(),
		}
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil
}

func (s *AuthService) Logout(userId string) *exception.Error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		s.log.Error("invalid user ID format", zap.String("userId", userId))
		return exception.ERROR_INVALID_USER_ID
	}

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		s.log.Error("failed to convert userId to ObjectID", zap.Error(err))
		return exception.ERROR_INVALID_USER_ID
	}

	_, err = s.tokenCollection.DeleteMany(ctx, bson.M{"user_id": objectId})
	if err != nil {
		s.log.Error("failed to delete tokens", zap.Error(err))
		return exception.ERROR_DELETE_TOKEN
	}

	return nil
}
