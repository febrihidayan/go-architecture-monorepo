package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/hashicorp/go-multierror"
)

type DeviceToken struct {
	ID        common.ID
	UserId    string
	Token     string
	OsName    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeviceTokenDto struct {
	ID        *common.ID
	UserId    string
	Token     string
	OsName    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeviceTokenQueryParams struct {
	UserId string
}

const (
	DeviceOsAndroid = "ANDROID"
	DeviceOsIOS     = "IOS"
)

func NewDeviceToken(x DeviceTokenDto, finds ...*DeviceToken) *DeviceToken {

	result := DeviceToken{
		ID:        common.NewID(),
		UserId:    x.UserId,
		Token:     x.Token,
		OsName:    x.OsName,
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

func (x *DeviceToken) Validate() (err *multierror.Error) {
	if x.UserId == "" {
		err = multierror.Append(err, lang.Trans("filled", "UserId"))
	}
	if x.Token == "" {
		err = multierror.Append(err, lang.Trans("filled", "Token"))
	}
	if x.OsName == "" {
		err = multierror.Append(err, lang.Trans("filled", "OsName"))
	}

	return
}
