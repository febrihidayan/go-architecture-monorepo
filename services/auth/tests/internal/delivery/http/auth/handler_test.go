package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	auth_http "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/auth"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/usecases"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthHandlerSuite struct {
	suite.Suite
	Route       *mux.Router
	Cfg         *config.AuthConfig
	Http        *auth_http.AuthHttpHandler
	Response    *httptest.ResponseRecorder
	AuthUsecase *usecases.AuthUsecaseMock
	Error       *exceptions.CustomError
	Token       string
	Id          common.ID
}

func (x *AuthHandlerSuite) SetupTest() {
	x.AuthUsecase = new(usecases.AuthUsecaseMock)

	x.Response = httptest.NewRecorder()
	x.Route = mux.NewRouter()

	x.Http = &auth_http.AuthHttpHandler{
		Cfg:         x.Cfg,
		AuthUsecase: x.AuthUsecase,
	}

	x.Error = &exceptions.CustomError{
		Status: exceptions.ERRREPOSITORY,
		Errors: multierror.Append(errors.New(mock.Anything)),
	}

	x.Token = "Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6InNpbTIifQ.eyJleHAiOjE3MDM4NjI3NTgsImp0aSI6Im1uYjIzdmNzcnQ3NTZ5dWlvbW5idmN4OThlcnR5dWlvcCIsInJvbGVzIjpbInN1cGVyYWRtaW5pc3RyYXRvciJdLCJzdWIiOiIzYzkyYzE4MC1mOGE0LTQ0ZGMtYWJlZS00ZDg2ODhmNDBjYWUifQ.aMKIC5uTqAxEEAhZ-aqvuhDYfc56M_6ZukUGImoxdLs"
	x.Id = common.NewID()

}

func TestAclHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}

type HandlerParams struct {
	method   string
	path     string
	payload  Any
	expected int
}

func BodyHelper(v any) *bytes.Buffer {
	b, _ := json.Marshal(v)
	return bytes.NewBuffer(b)
}
