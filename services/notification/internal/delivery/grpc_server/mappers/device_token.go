package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

func ToDomainDeviceTokenDto(x *notificationPb.DeviceToken) entities.DeviceTokenDto {
	id, _ := common.StringToID(x.GetId())
	return entities.DeviceTokenDto{
		ID:        &id,
		UserId:    x.GetUserId(),
		Token:     x.GetToken(),
		CreatedAt: x.CreatedAt.AsTime(),
		UpdatedAt: x.UpdatedAt.AsTime(),
	}
}
