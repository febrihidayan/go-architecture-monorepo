package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/hashicorp/go-multierror"
)

type Role struct {
	ID          common.ID
	Name        string
	DisplayName string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RoleDto struct {
	ID          *common.ID
	Name        string
	DisplayName string
	Description string
}

type RoleQueryParams struct {
	Search  string
	Page    int
	PerPage int
}

type RoleMeta struct {
	Data  []*Role
	Total int
}

func NewRole(x RoleDto) *Role {
	id := common.NewID()

	if x.ID != nil {
		id = *x.ID
	}

	return &Role{
		ID:          id,
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	}
}

func (x *Role) Validate() (err *multierror.Error) {
	if x.Name == "" {
		err = multierror.Append(err, lang.Trans("filled", "Name"))
	}
	if x.DisplayName == "" {
		err = multierror.Append(err, lang.Trans("filled", "DisplayName"))
	}

	return
}
