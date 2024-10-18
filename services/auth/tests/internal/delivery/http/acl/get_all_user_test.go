package acl

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func (x *AclHandlerSuite) TestGetAllUser() {
	testCases := []struct {
		name  string
		param HandlerParams
		mock  func(m HandlerParams)
	}{
		{
			name: "Success Positive Case",
			param: HandlerParams{
				method: http.MethodGet,
				path:   "/v1/auth/acl/user/{id}",
				payload: Any{
					x.Token,
					x.Id.String(),
				},
				expected: 200,
			},
			mock: func(m HandlerParams) {
				x.AclUsecase.On("GetAllUser", mock.Anything, mock.Anything).Return(&entities.AclMeta{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: http.MethodGet,
				path:   "/v1/auth/acl/user/{id}",
				payload: Any{
					"",
					x.Id.String(),
				},
				expected: 400,
			},
			mock: func(m HandlerParams) {
				x.AclUsecase.On("GetAllUser", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			id := tc.param.payload.Get(1).(string)

			req := httptest.NewRequest(tc.param.method, tc.param.path, nil)
			req.Header.Set("Authorization", token)
			re := mux.SetURLVars(req, map[string]string{"id": id})

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.GetAllUser).ServeHTTP(x.Response, re)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
