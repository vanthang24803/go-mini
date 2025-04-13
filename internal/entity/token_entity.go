package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id"`
	Token    string             `bson:"token"`
	Type     string             `bson:"type"`
	IP       string             `bson:"ip"`
	Checksum string             `bson:"checksum"`

	ExpiresAt time.Time `bson:"expires_at"`
	CreatedAt time.Time `bson:"created_at"`
}
