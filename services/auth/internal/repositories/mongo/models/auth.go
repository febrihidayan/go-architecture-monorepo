package models

import (
	"time"
)

type Auth struct {
	ID              string    `bson:"_id"`
	UserId          string    `bson:"user_id"`
	Email           string    `bson:"email"`
	Password        string    `bson:"password"`
	EmailVerifiedAt time.Time `bson:"email_verified_at"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}

func (Auth) TableName() string {
	return "auths"
}
