package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type NotificationListResponse struct {
	ID        common.ID   `json:"id"`
	Content   interface{} `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func MapNotificationListResponse(x *entities.Notification) NotificationListResponse {
	return NotificationListResponse{
		ID:        x.ID,
		Content:   x.GetData(),
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func MapNotificationListResponses(x []*entities.Notification) []NotificationListResponse {
	var result []NotificationListResponse
	for _, v := range x {
		result = append(result, MapNotificationListResponse(v))
	}
	return result
}
