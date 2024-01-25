package user

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/request"
	"github.com/stretchr/testify/mock"
)

func (x *UserHandlerSuite) TestGetAll() {
	testCases := []struct {
		name  string
		param HandlerParams
		mock  func(m HandlerParams)
	}{
		{
			name: "Success Positive Case",
			param: HandlerParams{
				method: http.MethodGet,
				path:   "/v1/users",
				payload: Any{
					x.Token,
				},
				expected: 200,
			},
			mock: func(m HandlerParams) {
				x.UserUsecase.On("GetAll", mock.Anything, mock.Anything).Return(&entities.UserMeta{}, nil)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			params := request.UserQueryParams{}
			token := tc.param.payload.Get(0).(string)

			req := httptest.NewRequest(tc.param.method, tc.param.path, nil)
			req.Header.Set("Authorization", token)

			if err := utils.MapQueryParams(req, &params); err != nil {
				fmt.Println("error params:", err)
				x.Fail(err.Error())
			}

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.GetAll).ServeHTTP(x.Response, req)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
