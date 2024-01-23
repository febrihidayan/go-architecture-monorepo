package role

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/request"
	"github.com/stretchr/testify/mock"
)

func (x *RoleHandlerSuite) TestCreate() {
	payload := request.RoleCreateRequest{
		Name:        "superadmin",
		DisplayName: "Super Admin",
		Description: "Super Admin",
	}

	testCases := []struct {
		name  string
		param HandlerParams
		mock  func(m HandlerParams)
	}{
		{
			name: "Success Positive Case",
			param: HandlerParams{
				method:   "POST",
				path:     "/v1/auth/role",
				payload:  Any{x.Token, payload},
				expected: 201,
			},
			mock: func(m HandlerParams) {
				x.RoleUsecase.On("Create", mock.Anything, mock.Anything).Return(&entities.Role{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method:   "POST",
				path:     "/v1/auth/role",
				payload:  Any{x.Token, request.RoleCreateRequest{}},
				expected: 422,
			},
			mock: func(m HandlerParams) {
				x.RoleUsecase.On("Create", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			body := tc.param.payload.Get(1).(request.RoleCreateRequest)

			req := httptest.NewRequest(tc.param.method, tc.param.path, BodyHelper(body))
			req.Header.Set("Authorization", token)

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.Create).ServeHTTP(x.Response, req)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
