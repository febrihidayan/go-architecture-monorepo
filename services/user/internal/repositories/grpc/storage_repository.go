package grpc_repositories

import (
	"context"

	storagePb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/storage"

	"google.golang.org/grpc"
)

type StorageRepository struct {
	svc storagePb.StorageServicesClient
	ctx context.Context
}

func NewStorageRepository(con *grpc.ClientConn) StorageRepository {
	client := storagePb.NewStorageServicesClient(con)

	return StorageRepository{
		svc: client,
		ctx: nil,
	}
}

func (x *StorageRepository) UpdateCloudApprove(ctx context.Context, url []string) error {
	_, err := x.svc.UpdateCloudApprove(ctx, &storagePb.CloudApproveRequest{
		Url: url,
	})

	return err
}

func (x *StorageRepository) DeleteCloudApprove(ctx context.Context, url []string) error {
	_, err := x.svc.DeleteCloudApprove(ctx, &storagePb.CloudApproveRequest{
		Url: url,
	})

	return err
}
