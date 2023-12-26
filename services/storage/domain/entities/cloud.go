package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/hashicorp/go-multierror"
)

const (
	CloudStatusPending = "pending"
	CloudStatusApprove = "approve"
)

type Cloud struct {
	ID        common.ID
	Name      string
	Url       string
	Status    string
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CloudDto struct {
	ID        *common.ID
	Name      string
	Url       string
	Status    string
	CreatedBy string
	File      File
}

type CloudQueryParams struct {
	Search    string
	Status    string
	CreatedAt time.Time
}

func NewCloud(x CloudDto, finds ...*Cloud) *Cloud {

	cloud := Cloud{
		ID:        common.NewID(),
		Name:      x.Name,
		Url:       x.Url,
		Status:    x.Status,
		CreatedBy: x.CreatedBy,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	if x.ID != nil {
		cloud.ID = *x.ID
	}

	for _, item := range finds {
		cloud.CreatedBy = item.CreatedBy
		cloud.CreatedAt = item.CreatedAt
	}

	return &cloud
}

func (x *Cloud) Validate() (err *multierror.Error) {
	if x.Name == "" {
		err = multierror.Append(err, lang.Trans("filled", "Name"))
	}
	if x.Url == "" {
		err = multierror.Append(err, lang.Trans("filled", "Url"))
	}
	if x.CreatedBy == "" {
		err = multierror.Append(err, lang.Trans("filled", "CreatedBy"))
	}

	return
}

func (x *Cloud) SetUrl(url string) {
	x.Url = url
}

func (x *Cloud) SetStatus(status string) {
	x.Status = status
}
