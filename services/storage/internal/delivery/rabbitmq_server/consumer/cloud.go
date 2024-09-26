package customer

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	storagePb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/storage"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (x *CustomerRabbitMQ) DeleteCloudApprove(ctx context.Context, body []byte) error {
	var req storagePb.CloudApproveRequest

	if err := rabbitmq.UnmarshalProto(body, &req); err != nil {
		log.Printf("DeleteCloudApprove#Failed to unmarshal: %v", err)
		return err
	}

	payloads := make([]*entities.Cloud, 0)

	for _, url := range req.GetUrl() {
		payloads = append(payloads, &entities.Cloud{
			Url:    url,
			Status: entities.CloudStatusApprove,
		})
	}

	if err := x.cloudUsecase.Deletes(ctx, payloads); err != nil {
		return status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return nil
}

func (x *CustomerRabbitMQ) UpdateCloudApprove(ctx context.Context, body []byte) error {
	var req storagePb.CloudApproveRequest

	if err := rabbitmq.UnmarshalProto(body, &req); err != nil {
		log.Printf("UpdateCloudApprove#Failed to unmarshal: %v", err)
		return err
	}

	payloads := make([]*entities.Cloud, 0)

	for _, url := range req.GetUrl() {
		payloads = append(payloads, &entities.Cloud{
			Url:    url,
			Status: entities.CloudStatusApprove,
		})
	}

	if err := x.cloudUsecase.UpdateStatus(ctx, payloads); err != nil {
		return status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return nil
}
