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
	"go.uber.org/zap"
)

type AuthService struct {
	userCollection *mongo.Collection
	log            *zap.Logger
}

func NewAuthService() *AuthService {
	return &AuthService{
		userCollection: database.GetCollection(constant.COLLECTION_USER),
		log:            logger.GetLogger(),
	}
}

func (s *AuthService) Login(req *dto.LoginRequest) (*entity.User, *exception.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entity.User

	err := s.userCollection.FindOne(ctx, bson.M{
		"email": req.Email,
	}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.ERROR_INVALID_CREDENTIAL
		}
		return nil, exception.ERROR_INVALID_CREDENTIAL
	}

	err = common.ComparePassword(req.Password, user.HashedPassword)

	if err != nil {
		return nil, exception.ERROR_INVALID_CREDENTIAL
	}

	return &user, nil
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*string, *exception.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser := &entity.User{}
	err := s.userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(existingUser)
	if err == nil {
		return nil, exception.ERROR_EMAIL_EXISTED
	}

	err = s.userCollection.FindOne(ctx, bson.M{"username": req.Username}).Decode(existingUser)
	if err == nil {
		return nil, exception.ERROR_CODE_USERNAME_EXISTED
	}
	if err != mongo.ErrNoDocuments {
		return nil, exception.ERROR_NO_DOCUMENT
	}

	hash, errHash := common.HashPassword(req.Password)
	if errHash != nil {
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
		return nil, &exception.Error{
			Code:    0,
			Message: err.Error(),
		}
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil

}
