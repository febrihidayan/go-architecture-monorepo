package models

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
)

type Permission struct {
	ID          common.ID `bson:"_id"`
	Name        string    `bson:"name"`
	DisplayName string    `bson:"display_name"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type PermissionMeta struct {
	Data    []*Permission `bson:"data,omitempty"`
	Page    int           `json:"page,omitempty"`
	PerPage int           `json:"per_page,omitempty"`
	Total   int           `json:"total,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
