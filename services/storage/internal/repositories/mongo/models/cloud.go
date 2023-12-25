package models

import (
	"time"
)

type Cloud struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Url       string    `bson:"url"`
	CreatedBy string    `bson:"created_by"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type CloudMeta struct {
	Data    []*Cloud `bson:"data,omitempty"`
	Page    int      `json:"page,omitempty"`
	PerPage int      `json:"per_page,omitempty"`
	Total   int      `json:"total,omitempty"`
}

func (Cloud) TableName() string {
	return "clouds"
}
