package role

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/request"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func (x *RoleHandlerSuite) TestUpdate() {
	payload := request.RoleUpdateRequest{
		ID:          &x.Id,
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
				method: "PUT",
				path:   "/v1/auth/role/{id}",
				payload: Any{
					x.Token,
					payload,
					x.Id.String(),
				},
				expected: 200,
			},
			mock: func(m HandlerParams) {
				x.RoleUsecase.On("Update", mock.Anything, mock.Anything).Return(&entities.Role{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: "PUT",
				path:   "/v1/auth/role/{id}",
				payload: Any{
					x.Token,
					request.RoleUpdateRequest{},
					x.Id.String(),
				},
				expected: 422,
			},
			mock: func(m HandlerParams) {
				x.RoleUsecase.On("Update", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			body := tc.param.payload.Get(1).(request.RoleUpdateRequest)
			id := tc.param.payload.Get(2).(string)

			req := httptest.NewRequest(tc.param.method, tc.param.path, BodyHelper(body))
			req.Header.Set("Authorization", token)
			re := mux.SetURLVars(req, map[string]string{"id": id})

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.Update).ServeHTTP(x.Response, re)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
