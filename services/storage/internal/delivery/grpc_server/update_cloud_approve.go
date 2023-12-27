package grpc_server

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	storagePb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/storage"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (x server) UpdateCloudApprove(ctx context.Context, req *storagePb.UpdateCloudApproveRequest) (*emptypb.Empty, error) {
	payloads := make([]*entities.Cloud, 0)

	for _, url := range req.GetUrl() {
		payloads = append(payloads, &entities.Cloud{
			Url:    url,
			Status: entities.CloudStatusApprove,
		})
	}

	if err := x.cloudUsecase.UpdateStatus(ctx, payloads); err != nil {
		return nil, status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return &emptypb.Empty{}, nil
}
