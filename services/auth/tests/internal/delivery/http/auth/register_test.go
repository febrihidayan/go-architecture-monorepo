package auth

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/request"
	"github.com/stretchr/testify/mock"
)

func (x *AuthHandlerSuite) TestRegister() {
	payload := request.AuthRegisterRequest{
		Name:            "Name",
		Email:           "name@example.com",
		Password:        "password",
		ConfirmPassword: "password",
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
				path:   "/v1/auth/register",
				payload: Any{
					x.Token,
					payload,
				},
				expected: 201,
			},
			mock: func(m HandlerParams) {
				x.AuthUsecase.On("Register", mock.Anything, mock.Anything).Return(&entities.AclMeta{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: http.MethodPost,
				path:   "/v1/auth/register",
				payload: Any{
					x.Token,
					request.AuthRegisterRequest{},
				},
				expected: 422,
			},
			mock: func(m HandlerParams) {
				x.AuthUsecase.On("Register", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			body := tc.param.payload.Get(1).(request.AuthRegisterRequest)

			req := httptest.NewRequest(tc.param.method, tc.param.path, BodyHelper(body))
			req.Header.Set("Authorization", token)

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.Register).ServeHTTP(x.Response, req)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
