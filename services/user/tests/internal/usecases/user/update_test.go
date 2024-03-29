package user

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *UserUsecaseSuite) TestUpdate() {
	id := common.NewID()
	var (
		user             *entities.User
		userDeleteAvatar *entities.User
		fullPath         = "https://testing.s3.ap-southeast-1.amazonaws.com/storage_test/upload.jpg"
	)

	payloadDto := entities.UserDto{
		ID:       &id,
		Name:     "Admin",
		Email:    "admin@app.com",
		Avatar:   fullPath,
		LangCode: entities.UserLangEN,
		Auth: entities.Auth{
			Password: "password",
		},
	}

	payloadDeleteAvatarDto := entities.UserDto{
		ID:       &id,
		Name:     "Admin",
		Email:    "admin@app.com",
		LangCode: entities.UserLangEN,
		Auth: entities.Auth{
			Password: "password",
		},
	}

	user = &entities.User{
		ID:        id,
		Name:      "Admin",
		Email:     "admin@app.com",
		Avatar:    fullPath,
		LangCode:  entities.UserLangEN,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	userDeleteAvatar = &entities.User{
		ID:        id,
		Name:      "Admin",
		Email:     "admin@app.com",
		LangCode:  entities.UserLangEN,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	auth := entities.Auth{
		UserId:   id.String(),
		Email:    user.Email,
		Password: payloadDto.Auth.Password,
	}

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Find", payloadDto.ID.String()).Return(user, nil)

				x.authRepo.Mock.On("CreateOrUpdate", &auth).Return(nil)

				x.userRepo.Mock.On("Update", user).Return(nil)

				x.storageGrpcRepo.Mock.On("UpdateCloudApprove", []string{fullPath}).Return(nil)

				result, err := x.userUsecase.Update(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(user, result)
			},
		},
		{
			name: "Success Upload Avatar Positive Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Find", payloadDto.ID.String()).Return(userDeleteAvatar, nil)

				x.authRepo.Mock.On("CreateOrUpdate", &auth).Return(nil)

				x.userRepo.Mock.On("Update", user).Return(nil)

				x.storageGrpcRepo.Mock.On("UpdateCloudApprove", []string{fullPath}).Return(nil)

				result, err := x.userUsecase.Update(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(user, result)
			},
		},
		{
			name: "Success Delete Avatar Positive Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Find", payloadDto.ID.String()).Return(user, nil)

				x.authRepo.Mock.On("CreateOrUpdate", &auth).Return(nil)

				x.userRepo.Mock.On("Update", userDeleteAvatar).Return(nil)

				x.storageGrpcRepo.Mock.On("DeleteCloudApprove", []string{fullPath}).Return(nil)

				result, err := x.userUsecase.Update(context.Background(), payloadDeleteAvatarDto)
				x.Nil(err)
				x.Equal(userDeleteAvatar, result)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Find", payloadDto.ID.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.userUsecase.Update(context.Background(), payloadDto)
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
