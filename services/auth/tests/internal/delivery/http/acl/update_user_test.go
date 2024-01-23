package acl

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/request"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func (x *AclHandlerSuite) TestUpdateUser() {
	payload := request.AclUserUpdateRequest{
		Permissions: []string{
			x.Id.String(),
		},
		Roles: []string{
			x.Id.String(),
		},
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
				path:   "/v1/auth/acl/user/{id}",
				payload: Any{
					x.Token,
					x.Id.String(),
					payload,
				},
				expected: 200,
			},
			mock: func(m HandlerParams) {
				x.AclUsecase.On("UpdateUser", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: http.MethodPut,
				path:   "/v1/auth/acl/user/{id}",
				payload: Any{
					x.Token,
					"",
					payload,
				},
				expected: 400,
			},
			mock: func(m HandlerParams) {
				x.AclUsecase.On("UpdateUser", mock.Anything, mock.Anything).Return(x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			id := tc.param.payload.Get(1).(string)
			body := tc.param.payload.Get(2).(request.AclUserUpdateRequest)

			req := httptest.NewRequest(tc.param.method, tc.param.path, BodyHelper(body))
			req.Header.Set("Authorization", token)
			re := mux.SetURLVars(req, map[string]string{"id": id})

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.UpdateUser).ServeHTTP(x.Response, re)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
