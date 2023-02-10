package models

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
)

type User struct {
	ID        common.ID `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Role      string    `bson:"role"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
