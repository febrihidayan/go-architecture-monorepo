package cloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	cloud_http "github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/http/delivery/cloud"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/tests/mocks/usecases"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CloudHandlerSuite struct {
	suite.Suite
	Route        *mux.Router
	Cfg          *config.StorageConfig
	Http         *cloud_http.CloudHttpHandler
	Response     *httptest.ResponseRecorder
	CloudUsecase *usecases.CloudUsecaseMock
	Error        *exceptions.CustomError
	Token        string
	Id           common.ID
}

func (x *CloudHandlerSuite) SetupTest() {
	x.Cfg = &config.StorageConfig{
		MaxUpload: 15,
	}

	x.CloudUsecase = new(usecases.CloudUsecaseMock)

	x.Response = httptest.NewRecorder()
	x.Route = mux.NewRouter()

	x.Http = &cloud_http.CloudHttpHandler{
		Cfg:          x.Cfg,
		CloudUsecase: x.CloudUsecase,
	}

	x.Error = &exceptions.CustomError{
		Status: exceptions.ERRREPOSITORY,
		Errors: multierror.Append(errors.New(mock.Anything)),
	}

	x.Token = "Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6InNpbTIifQ.eyJleHAiOjE3MDM4NjI3NTgsImp0aSI6Im1uYjIzdmNzcnQ3NTZ5dWlvbW5idmN4OThlcnR5dWlvcCIsInJvbGVzIjpbInN1cGVyYWRtaW5pc3RyYXRvciJdLCJzdWIiOiIzYzkyYzE4MC1mOGE0LTQ0ZGMtYWJlZS00ZDg2ODhmNDBjYWUifQ.aMKIC5uTqAxEEAhZ-aqvuhDYfc56M_6ZukUGImoxdLs"
	x.Id = common.NewID()

}

func TestAclHandler(t *testing.T) {
	suite.Run(t, new(CloudHandlerSuite))
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
