package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type UserMeta struct {
	Data    []*User `bson:"data,omitempty"`
	Page    int     `json:"page,omitempty"`
	PerPage int     `json:"per_page,omitempty"`
	Total   int     `json:"total,omitempty"`
}

func (User) TableName() string {
	return "users"
}
