package repositories

import (
	"context"
)

type StorageRepository interface {
	UpdateCloudApprove(ctx context.Context, url []string) error
}
