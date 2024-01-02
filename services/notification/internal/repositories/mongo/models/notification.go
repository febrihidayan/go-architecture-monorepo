package models

import (
	"time"
)

type Notification struct {
	ID        string    `bson:"_id"`
	UserId    string    `bson:"user_id"`
	Type      string    `bson:"type"`
	Data      string    `bson:"data"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type NotificationMeta struct {
	Data    []*Notification `bson:"data,omitempty"`
	Page    int             `json:"page,omitempty"`
	PerPage int             `json:"per_page,omitempty"`
	Total   int             `json:"total,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}
