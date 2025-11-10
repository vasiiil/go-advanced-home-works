package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error any `json:"error"`
}

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func JsonError(w http.ResponseWriter, err any, statusCode int) {
	resp := ErrorResponse{
		Error: err,
	}
	Json(w, resp, statusCode)
}

func BadRequestJson(w http.ResponseWriter, data any) {
	JsonError(w, data, http.StatusBadRequest)
}
