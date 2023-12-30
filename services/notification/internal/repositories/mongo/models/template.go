package models

import (
	"time"
)

type Template struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Data      string    `bson:"data"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type TemplateMeta struct {
	Data    []*Template `bson:"data,omitempty"`
	Page    int         `json:"page,omitempty"`
	PerPage int         `json:"per_page,omitempty"`
	Total   int         `json:"total,omitempty"`
}

func (Template) TableName() string {
	return "templates"
}
