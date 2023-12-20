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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDto struct {
	ID    *common.ID
	Name  string
	Email string
}

func NewUser(x UserDto) *User {
	id := common.NewID()

	if x.ID != nil {
		id = *x.ID
	}

	return &User{
		ID:        id,
		Name:      x.Name,
		Email:     x.Email,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}
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
