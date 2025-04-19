package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username       string             `bson:"username" json:"username"`
	Email          string             `bson:"email" json:"email"`
	HashedPassword string             `bson:"hashed_password" json:"-"`
	Phone          string             `bson:"phone" json:"phone"`
	FirstName      string             `bson:"first_name" json:"first_name"`
	LastName       string             `bson:"last_name" json:"last_name"`
	Avatar         string             `bson:"avatar" json:"avatar"`
	Gender         string             `bson:"gender" json:"gender"`
	Address        string             `bson:"address" json:"address"`
	DateOfBirth    time.Time          `bson:"date_of_birth" json:"date_of_birth"`

	Roles    []string `bson:"role" json:"roles"`
	Active   bool     `bson:"active" json:"active"`
	Version  int      `bson:"version" json:"version"`
	Timezone string   `bson:"timezone" json:"timezone"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
