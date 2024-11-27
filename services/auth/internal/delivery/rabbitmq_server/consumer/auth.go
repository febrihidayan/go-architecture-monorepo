package customer

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	authPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (x *CustomerRabbitMQ) AuthDelete(ctx context.Context, body []byte) error {
	var req authPb.FindByIdRequest

	if err := rabbitmq.UnmarshalProto(body, &req); err != nil {
		log.Printf("AuthDelete#Failed to unmarshal: %v", err)
		return err
	}

	if err := x.authUsecase.DeleteByUserID(ctx, req.GetId()); err != nil {
		return status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return nil
}
