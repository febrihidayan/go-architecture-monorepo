package cloud

import (
	"context"
	"errors"
	"os"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *CloudUsecaseSuite) TestCreate() {
	id := common.NewID()
	var (
		cloud    *entities.Cloud
		fullPath = "https://testing.s3.ap-southeast-1.amazonaws.com/storage_test/upload.jpg"
	)

	fileName, _ := os.CreateTemp("", "upload.jpg")
	payloadDto := entities.CloudDto{
		ID:        &id,
		Name:      "upload",
		CreatedBy: id.String(),
		File: entities.File{
			Name: fileName.Name(),
		},
	}

	cloud = &entities.Cloud{
		ID:        id,
		Name:      "upload",
		Url:       fullPath,
		Status:    entities.CloudStatusPending,
		CreatedBy: id.String(),
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.awsService.Mock.On("UploadFile", &payloadDto.File).Return(fullPath, nil)

				x.cloudRepo.Mock.On("Create", cloud).Return(nil)

				result, err := x.cloudUsecase.Create(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(cloud, result)
			},
		},
		{
			name: "Failed Upload Negatif Case",
			tests: func(arg Any) {
				x.awsService.Mock.On("UploadFile", &payloadDto.File).Return("", errors.New(mock.Anything))

				_, err := x.cloudUsecase.Create(context.Background(), payloadDto)
				e := &exceptions.CustomError{
					Status: exceptions.ERRBUSSINESS,
					Errors: multierror.Append(errors.New(mock.Anything)),
				}

				x.Equal(err, e)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.awsService.Mock.On("UploadFile", &payloadDto.File).Return(fullPath, nil)

				x.cloudRepo.Mock.On("Create", cloud).Return(errors.New(mock.Anything))

				_, err := x.cloudUsecase.Create(context.Background(), payloadDto)
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
