package acl

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *AclHandlerSuite) TestAccessUserLogin() {
	testCases := []struct {
		name  string
		param HandlerParams
		mock  func(m HandlerParams)
	}{
		{
			name: "Success Positive Case",
			param: HandlerParams{
				method: "GET",
				path:   "/v1/auth/acl/access",
				payload: Any{
					x.Token,
				},
				expected: 200,
			},
			mock: func(m HandlerParams) {
				x.AclUsecase.On("AccessUserLogin", mock.Anything, mock.Anything).Return(&entities.AclMeta{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: "GET",
				path:   "/v1/auth/acl/access",
				payload: Any{
					"",
				},
				expected: 400,
			},
			mock: func(m HandlerParams) {
				x.AclUsecase.On("AccessUserLogin", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)

			req := httptest.NewRequest(tc.param.method, tc.param.path, nil)
			req.Header.Set("Authorization", token)

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.AccessUserLogin).ServeHTTP(x.Response, req)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
