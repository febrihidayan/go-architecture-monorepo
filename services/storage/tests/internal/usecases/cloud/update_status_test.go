package cloud

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *CloudUsecaseSuite) TestUpdateStatus() {
	id := common.NewID()
	var (
		cloud    *entities.Cloud
		payloads []*entities.Cloud
		fullPath = "https://testing.s3.ap-southeast-1.amazonaws.com/storage_test/upload.jpg"
	)

	cloud = &entities.Cloud{
		ID:        id,
		Name:      "upload",
		Url:       fullPath,
		Status:    entities.CloudStatusPending,
		CreatedBy: id.String(),
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	payloads = append(payloads, &entities.Cloud{
		ID:     id,
		Status: entities.CloudStatusApprove,
	})

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.cloudRepo.Mock.On("Find", cloud.ID.String()).Return(cloud, nil)

				x.cloudRepo.Mock.On("Update", cloud).Return(nil)

				err := x.cloudUsecase.UpdateStatus(context.Background(), payloads)
				x.Nil(err)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				// reset status pending
				cloud.SetStatus(entities.CloudStatusPending)
				x.cloudRepo.Mock.On("Find", cloud.ID.String()).Return(cloud, nil)

				x.cloudRepo.Mock.On("Update", cloud).Return(errors.New(mock.Anything))

				err := x.cloudUsecase.UpdateStatus(context.Background(), payloads)
				e := &exceptions.CustomError{
					Status: exceptions.ERRSYSTEM,
					Errors: multierror.Append(errors.New(mock.Anything)),
				}

				x.Equal(err, e)
			},
		},
	}

	for _, arg := range args {
		x.Run(arg.name, func() {
			x.SetupTest()
			arg.tests(arg.args)
		})
	}
}
