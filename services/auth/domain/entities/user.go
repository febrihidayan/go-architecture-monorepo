package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
)

type User struct {
	ID              common.ID
	Name            string
	Email           string
	Role            string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
