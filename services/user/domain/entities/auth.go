package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
)

type Auth struct {
	ID        common.ID
	UserId    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
