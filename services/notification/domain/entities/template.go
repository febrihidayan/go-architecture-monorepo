package entities

import (
	"encoding/json"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/hashicorp/go-multierror"
)

type Template struct {
	ID        common.ID
	Name      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateDto struct {
	ID        *common.ID
	Name      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateQueryParams struct {
	Search string
}

func NewTemplate(x TemplateDto, finds ...*Template) *Template {

	result := Template{
		ID:        common.NewID(),
		Name:      x.Name,
		Data:      x.Data,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	if x.ID != nil {
		result.ID = *x.ID
	}

	for _, item := range finds {
		result.CreatedAt = item.CreatedAt
	}

	return &result
}

func (x *Template) Validate() (err *multierror.Error) {
	if x.Name == "" {
		err = multierror.Append(err, lang.Trans("filled", "Name"))
	}
	if x.Data == "" {
		err = multierror.Append(err, lang.Trans("filled", "Data"))
	}

	return
}

func (x *Template) GetData() (data interface{}) {
	json.Unmarshal([]byte(x.Data), &data)

	return
}
