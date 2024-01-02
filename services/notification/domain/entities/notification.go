package entities

import (
	"encoding/json"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/hashicorp/go-multierror"
)

type Notification struct {
	ID        common.ID
	UserId    string
	Type      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NotificationDto struct {
	ID        *common.ID
	UserId    string
	Type      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NotificationQueryParams struct {
	UserId  string
	Page    int
	PerPage int
}

type NotificationMeta struct {
	Data    []*Notification
	Total   int
	Page    int
	PerPage int
}

func NewNotification(x NotificationDto, finds ...*Notification) *Notification {

	result := Notification{
		ID:        common.NewID(),
		UserId:    x.UserId,
		Type:      x.Type,
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

func (x *Notification) Validate() (err *multierror.Error) {
	if x.UserId == "" {
		err = multierror.Append(err, lang.Trans("filled", "UserId"))
	}
	if x.Type == "" {
		err = multierror.Append(err, lang.Trans("filled", "Type"))
	}
	if x.Data == "" {
		err = multierror.Append(err, lang.Trans("filled", "Data"))
	}

	return
}

func (x *Notification) SetData(data interface{}) {
	jsonData, _ := json.Marshal(data)

	x.Data = string(jsonData)
}

func (x *Notification) GetData() (data interface{}) {
	json.Unmarshal([]byte(x.Data), &data)

	return
}
