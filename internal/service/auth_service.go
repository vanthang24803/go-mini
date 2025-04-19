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
	ctx             context.Context
	userCollection  *mongo.Collection
	tokenCollection *mongo.Collection
	log             *zap.Logger
	redis           database.Redis
}

func NewAuthService() *AuthService {
	return &AuthService{
		ctx:             context.Background(),
		userCollection:  database.GetCollection(constant.COLLECTION_USER),
		tokenCollection: database.GetCollection(constant.COLLECTION_TOKEN),
		log:             logger.GetLogger(),
		redis:           database.NewRedisService(),
	}
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.TokenResponse, *exception.Error) {
	var cachedToken *dto.TokenResponse
	if err := s.redis.Get(s.ctx, req.Email, &cachedToken); err == nil && cachedToken != nil {
		return cachedToken, nil
	}

	var user entity.User
	if err := s.userCollection.FindOne(s.ctx, bson.M{"email": req.Email}).Decode(&user); err != nil {
		s.log.Error("failed to find user by email", zap.Error(err))
		return nil, exception.ERROR_INVALID_CREDENTIAL
	}

	if err := common.ComparePassword(req.Password, user.HashedPassword); err != nil {
		s.log.Error("invalid password", zap.Error(err))
		return nil, exception.ERROR_INVALID_CREDENTIAL
	}

	token, refreshToken, err := common.GenerateJWT(user.ID.Hex(), user.Username)
	if err != nil {
		s.log.Error("failed to generate JWT", zap.Error(err))
		return nil, exception.ERROR_CODE_INTERNAL_ERROR
	}

	var currentToken entity.Token
	err = s.tokenCollection.FindOne(s.ctx, bson.M{
		"user_id": user.ID,
		"type":    constant.REFRESH_TOKEN,
	}, options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})).Decode(&currentToken)

	finalRefreshToken := refreshToken
	if err == nil && !common.CompareWithNow(currentToken.ExpiresAt) {
		finalRefreshToken = currentToken.Token
	} else {
		newToken := entity.Token{
			UserID:    user.ID,
			Token:     refreshToken,
			Type:      constant.REFRESH_TOKEN,
			ExpiresAt: time.Now().Add(constant.REFRESH_TOKEN_EXPIRATION),
			CreatedAt: time.Now(),
			Checksum:  common.GenerateCheckSum(&user),
		}

		if _, err := s.tokenCollection.InsertOne(s.ctx, newToken); err != nil {
			s.log.Error("failed to insert token", zap.Error(err))
			return nil, exception.ERROR_INSERT_TOKEN
		}
	}

	response := &dto.TokenResponse{
		Type:         "Bearer",
		Exp:          time.Now().Add(7 * 24 * time.Hour).Unix(),
		Token:        token,
		RefreshToken: finalRefreshToken,
	}

	if err := s.redis.Set(s.ctx, req.Email, response, constant.ACCESS_TOKEN_EXPIRATION-time.Minute*5); err != nil {
		s.log.Error("failed to cache token", zap.Error(err))
	}

	return response, nil
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*string, *exception.Error) {

	existingUser := &entity.User{}
	err := s.userCollection.FindOne(s.ctx, bson.M{"email": req.Email}).Decode(existingUser)
	if err == nil {
		s.log.Error("email already exists", zap.String("email", req.Email))
		return nil, exception.ERROR_EMAIL_EXISTED
	}

	err = s.userCollection.FindOne(s.ctx, bson.M{"username": req.Username}).Decode(existingUser)
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
		Roles:          []string{constant.ROLE_ROOT},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Timezone:       "Asia/Ho_Chi_Minh",
		Version:        constant.VERSION,
	}

	result, err := s.userCollection.InsertOne(s.ctx, user)
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

	_, err = s.tokenCollection.DeleteMany(s.ctx, bson.M{"user_id": objectId})
	if err != nil {
		s.log.Error("failed to delete tokens", zap.Error(err))
		return exception.ERROR_DELETE_TOKEN
	}

	return nil
}
