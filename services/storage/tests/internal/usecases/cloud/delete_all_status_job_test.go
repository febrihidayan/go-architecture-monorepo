package cloud

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *CloudUsecaseSuite) TestDeleteAllStatusJob() {
	id := common.NewID()
	var (
		cloud    *entities.Cloud
		clouds   []*entities.Cloud
		fullPath = "https://testing.s3.ap-southeast-1.amazonaws.com/storage_test/upload.jpg"
	)

	params := entities.CloudQueryParams{
		Status: entities.CloudStatusPending,
	}

	cloud = &entities.Cloud{
		ID:        id,
		Name:      "upload",
		Url:       fullPath,
		CreatedBy: id.String(),
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	clouds = append(clouds, cloud)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.cloudRepo.Mock.On("All", &params).Return(clouds, nil)

				x.cloudRepo.Mock.On("Delete", cloud.ID.String()).Return(nil)

				err := x.cloudUsecase.DeleteAllStatusJob(context.Background())
				x.Nil(err)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.cloudRepo.Mock.On("All", &params).Return(clouds, nil)

				x.cloudRepo.Mock.On("Delete", cloud.ID.String()).Return(errors.New(mock.Anything))

				err := x.cloudUsecase.DeleteAllStatusJob(context.Background())
				x.NotNil(err)
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
