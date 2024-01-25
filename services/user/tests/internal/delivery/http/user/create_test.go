package user

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/request"
	"github.com/stretchr/testify/mock"
)

func (x *UserHandlerSuite) TestCreate() {
	payload := request.UserCreateRequest{
		Name:     "Name",
		Email:    "name@example.com",
		Password: "password",
	}

	testCases := []struct {
		name  string
		param HandlerParams
		mock  func(m HandlerParams)
	}{
		{
			name: "Success Positive Case",
			param: HandlerParams{
				method: http.MethodPost,
				path:   "/v1/user",
				payload: Any{
					x.Token,
					payload,
				},
				expected: 201,
			},
			mock: func(m HandlerParams) {
				x.UserUsecase.On("Create", mock.Anything, mock.Anything).Return(&entities.User{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: http.MethodPost,
				path:   "/v1/user",
				payload: Any{
					x.Token,
					request.UserCreateRequest{},
				},
				expected: 422,
			},
			mock: func(m HandlerParams) {
				x.UserUsecase.On("Create", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			body := tc.param.payload.Get(1).(request.UserCreateRequest)

			req := httptest.NewRequest(tc.param.method, tc.param.path, BodyHelper(body))
			req.Header.Set("Authorization", token)

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.Create).ServeHTTP(x.Response, req)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
