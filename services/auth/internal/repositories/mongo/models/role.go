package models

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
)

type Role struct {
	ID          common.ID `bson:"_id"`
	Name        string    `bson:"name"`
	DisplayName string    `bson:"display_name"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func (Role) TableName() string {
	return "roles"
}
