package response

import (
	"moonlogs/internal/lib/serialize"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta"`
}

type Meta struct {
	Page  int `json:"page"`
	Count int `json:"count"`
	Pages int `json:"pages"`
}

func Return(w http.ResponseWriter, success bool, code int, err error, data interface{}, meta Meta) {
	response := Response{
		Success: success,
		Code:    code,
		Error:   "",
		Data:    data,
		Meta:    meta,
	}

	if err != nil {
		response.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	err = serialize.NewJSONEncoder(w).Encode(response)
	if err != nil {
		response.Error = err.Error()
	}
}

func ReturnPlain(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	_ = serialize.NewJSONEncoder(w).Encode(data)
}
