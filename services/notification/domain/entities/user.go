package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
)

type User struct {
	ID        common.ID
	Name      string
	Email     string
	Avatar    string
	LangCode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// default english if lang code is empty
func (x *User) LangDefault() {
	if x.LangCode == "" {
		x.LangCode = TemplateLangEN
	}
}
