package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func Return(w http.ResponseWriter, success bool, code int, err error, data interface{}) {
	response := Response{
		Success: success,
		Code:    code,
		Error:   "",
		Data:    data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		response.Error = err.Error()
	}
}
