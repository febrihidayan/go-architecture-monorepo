package auth

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func (x *AuthHandlerSuite) TestEmailVerified() {
	testCases := []struct {
		name  string
		param HandlerParams
		mock  func(m HandlerParams)
	}{
		{
			name: "Success Positive Case",
			param: HandlerParams{
				method: http.MethodGet,
				path:   "/v1/auth/email/{token}",
				payload: Any{
					"example-token",
				},
				expected: 200,
			},
			mock: func(m HandlerParams) {
				x.AuthUsecase.On("EmailVerified", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: http.MethodGet,
				path:   "/v1/auth/email/{token}",
				payload: Any{
					"",
				},
				expected: 400,
			},
			mock: func(m HandlerParams) {
				x.AuthUsecase.On("EmailVerified", mock.Anything, mock.Anything).Return(x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)

			req := httptest.NewRequest(tc.param.method, tc.param.path, nil)
			re := mux.SetURLVars(req, map[string]string{"token": token})

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.EmailVerified).ServeHTTP(x.Response, re)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
