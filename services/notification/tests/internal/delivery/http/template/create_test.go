package template

import (
	"net/http"
	"net/http/httptest"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/request"
	"github.com/stretchr/testify/mock"
)

func (x *TemplateHandlerSuite) TestCreate() {
	payload := request.TemplateCreateRequest{
		Name: "welcome",
		Data: map[string]interface{}{
			"title": map[string]string{
				"id": "Selamat Datang",
			},
			"body": map[string]string{
				"id": "Hai {{.name}}, kamu telah menjadi bagian dari kami.",
			},
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
				method: http.MethodPost,
				path:   "/v1/notification/template",
				payload: Any{
					x.Token,
					payload,
				},
				expected: 201,
			},
			mock: func(m HandlerParams) {
				x.TemplateUsecase.On("Create", mock.Anything, mock.Anything).Return(&entities.Template{}, nil)
			},
		},
		{
			name: "Failed Negatif Case",
			param: HandlerParams{
				method: http.MethodPost,
				path:   "/v1/notification/template",
				payload: Any{
					x.Token,
					request.TemplateCreateRequest{
						Data: "",
					},
				},
				expected: 422,
			},
			mock: func(m HandlerParams) {
				x.TemplateUsecase.On("Create", mock.Anything, mock.Anything).Return(nil, x.Error)
			},
		},
	}

	for _, tc := range testCases {
		x.Run(tc.name, func() {
			x.SetupTest()

			token := tc.param.payload.Get(0).(string)
			body := tc.param.payload.Get(1).(request.TemplateCreateRequest)

			req := httptest.NewRequest(tc.param.method, tc.param.path, BodyHelper(body))
			req.Header.Set("Authorization", token)

			tc.mock(tc.param)
			http.HandlerFunc(x.Http.Create).ServeHTTP(x.Response, req)
			x.Equal(tc.param.expected, x.Response.Result().StatusCode)
		})
	}
}
