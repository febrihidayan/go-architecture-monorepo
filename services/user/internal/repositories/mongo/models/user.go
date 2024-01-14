package models

import (
	"time"
)

type User struct {
	ID              string    `bson:"_id"`
	Name            string    `bson:"name"`
	Email           string    `bson:"email"`
	Avatar          string    `bson:"avatar"`
	LangCode        string    `bson:"lang_code"`
	EmailVerifiedAt time.Time `bson:"email_verified_at"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
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
