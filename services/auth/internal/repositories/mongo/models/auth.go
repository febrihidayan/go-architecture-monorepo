package models

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
)

type Auth struct {
	ID        common.ID `bson:"_id"`
	UserId    string    `bson:"user_id"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Role      string    `bson:"role"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (Auth) TableName() string {
	return "auths"
}
