package validator

import (
	"encoding/json"
	"net/http"

	"github.com/kmacute/golvalidator"
)

func Make(s interface{}) interface{} {
	if errors := golvalidator.ValidateStructs(s); len(errors) > 0 {
		return errors
	}

	return nil
}

func ErrorJson(w http.ResponseWriter, code int, errors interface{}) {
	err := map[string]interface{}{
		"code":    code,
		"message": "Error",
		"errors":  errors,
	}
	w.WriteHeader(code)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}
