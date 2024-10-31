package user

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *UserUsecaseSuite) TestDelete() {
	id := common.NewID()

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Delete", id.String()).Return(nil)

				err := x.userUsecase.Delete(context.Background(), id.String())
				x.Nil(err)
			},
		},
		{
			name: "Failed Negative Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Delete", id.String()).Return(errors.New(mock.Anything))

				err := x.userUsecase.Delete(context.Background(), id.String())

				e := &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
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
