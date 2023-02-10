package utils

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

type jsonErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func RespondWithError(w http.ResponseWriter, code int, errors []error) {
	var errStr []string
	for _, err := range errors {
		errStr = append(errStr, err.Error())
	}
	jsonResponse := jsonErrorResponse{
		Code:    code,
		Message: "Error",
		Errors:  errStr,
	}
	response, _ := json.Marshal(jsonResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(jsonResponse{
		Code:    code,
		Message: "Success",
		Data:    payload,
		Meta:    nil,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseWithJsonMeta(w http.ResponseWriter, code int, payload interface{}, meta interface{}) {
	response, _ := json.Marshal(jsonResponse{
		Code:    code,
		Message: "Success",
		Data:    payload,
		Meta:    meta,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
