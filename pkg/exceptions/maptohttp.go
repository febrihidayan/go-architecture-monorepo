package exceptions

import (
	"net/http"
)

func MapToHttpStatusCode(status Status) int {
	var httpStatusCode int
	switch status {
	case ERRDOMAIN:
		httpStatusCode = http.StatusBadRequest
	case ERRBUSSINESS:
		httpStatusCode = http.StatusBadRequest
	case ERRSYSTEM:
		httpStatusCode = http.StatusInternalServerError
	case ERRNOTFOUND:
		httpStatusCode = http.StatusNotFound
	case ERRREPOSITORY:
		httpStatusCode = http.StatusInternalServerError
	case ERRUNKNOWN:
		httpStatusCode = http.StatusInternalServerError
	case ERRAUTHORIZED:
		httpStatusCode = http.StatusUnauthorized
	case ERRFORBIDDEN:
		httpStatusCode = http.StatusForbidden
	default:
		httpStatusCode = http.StatusInternalServerError
	}

	return httpStatusCode
}
