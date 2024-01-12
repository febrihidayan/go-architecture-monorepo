package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type DeviceTokenListResponse struct {
	ID        common.ID `json:"id"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token"`
	OsName    string    `json:"os_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapDeviceTokenListResponse(x *entities.DeviceToken) DeviceTokenListResponse {
	return DeviceTokenListResponse{
		ID:        x.ID,
		UserId:    x.UserId,
		Token:     x.Token,
		OsName:    x.OsName,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func MapDeviceTokenListResponses(x []*entities.DeviceToken) []DeviceTokenListResponse {
	var result []DeviceTokenListResponse
	for _, v := range x {
		result = append(result, MapDeviceTokenListResponse(v))
	}
	return result
}
