package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/hashicorp/go-multierror"
)

type Permission struct {
	ID          common.ID
	Name        string
	DisplayName string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PermissionDto struct {
	ID          *common.ID
	Name        string
	DisplayName string
	Description string
}

type PermissionQueryParams struct {
	Search  string
	Page    int
	PerPage int
}

type PermissionMeta struct {
	Data  []*Permission
	Total int
}

func NewPermission(x PermissionDto, finds ...*Permission) *Permission {
	permission := Permission{
		ID:          common.NewID(),
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	}

	if x.ID != nil {
		permission.ID = *x.ID
	}

	for _, item := range finds {
		permission.CreatedAt = item.CreatedAt
	}

	return &permission
}

func (x *Permission) Validate() (err *multierror.Error) {
	if x.Name == "" {
		err = multierror.Append(err, lang.Trans("filled", "Name"))
	}
	if x.DisplayName == "" {
		err = multierror.Append(err, lang.Trans("filled", "DisplayName"))
	}

	return
}
