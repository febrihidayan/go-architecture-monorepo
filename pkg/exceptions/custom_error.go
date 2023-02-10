package exceptions

import "github.com/hashicorp/go-multierror"

const (
	ERRDOMAIN     Status = 1
	ERRBUSSINESS  Status = 2
	ERRSYSTEM     Status = 3
	ERRNOTFOUND   Status = 4
	ERRREPOSITORY Status = 5
	ERRUNKNOWN    Status = 6
	ERRAUTHORIZED Status = 7
	ERRFORBIDDEN  Status = 8
)

type Status int

type CustomError struct {
	Status Status
	Errors *multierror.Error
}
