package models

import (
	"time"
)

type DeviceToken struct {
	ID        string    `bson:"_id"`
	UserId    string    `bson:"user_id"`
	Token     string    `bson:"token"`
	OsName    string    `bson:"os_name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (DeviceToken) TableName() string {
	return "device_tokens"
}
