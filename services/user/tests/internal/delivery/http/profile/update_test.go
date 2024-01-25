package profile

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/request"
	"github.com/stretchr/testify/mock"
)

func (x *ProfileHandlerSuite) TestUpdate() {
	payload := request.UserUpdateRequest{
		ID:       &x.Id,
		Name:     "Name",
		Email:    "name@example.com",
		LangCode: "id",
	}

	testCases := []struct {
		name  string
		param HandlerParams
		mock  func(m HandlerParams)
	}{
		{
			name: "Success Positive Case",
			param: HandlerParams{
				method: http.MethodPut,
				path:   "/v1/user/profile",
				payload: Any{
					x.Token,
					payload,
				},
				expected: 200,
			},
			mock: func(m HandlerParams) {
				x.ProfileUsecase.On("Update", mock.Anything, mock.Anything).Return(&entities.User{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: http.MethodPut,
				path:   "/v1/user/profile",
				payload: Any{
					x.Token,
					request.UserUpdateRequest{},
				},
				expected: 422,
			},
			mock: func(m HandlerParams) {
				x.ProfileUsecase.On("Update", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			body := tc.param.payload.Get(1).(request.UserUpdateRequest)

			req := httptest.NewRequest(tc.param.method, tc.param.path, BodyHelper(body))
			req.Header.Set("Authorization", token)

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.Update).ServeHTTP(x.Response, req)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
