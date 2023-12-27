package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"

	"github.com/hashicorp/go-multierror"
)

type User struct {
	ID        common.ID
	Name      string
	Email     string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDto struct {
	ID     *common.ID
	Name   string
	Email  string
	Avatar string
	Auth   Auth
}

type UserQueryParams struct {
	Search  string
	Page    int
	PerPage int
}

type UserMeta struct {
	Data  []*User
	Total int
}

func NewUser(x UserDto, finds ...*User) *User {
	user := User{
		ID:        common.NewID(),
		Name:      x.Name,
		Email:     x.Email,
		Avatar:    x.Avatar,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	if x.ID != nil {
		user.ID = *x.ID
	}

	for _, item := range finds {
		user.CreatedAt = item.CreatedAt
	}

	return &user
}

func (x *User) Validate() (err *multierror.Error) {
	if x.Name == "" {
		err = multierror.Append(err, lang.ErrNameRequired)
	}
	if x.Email == "" {
		err = multierror.Append(err, lang.ErrEmailRequired)
	}

	return
}
