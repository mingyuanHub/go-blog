package core

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type ResCode struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func JsonCode(w http.ResponseWriter, code int, data interface{}) {
	res := ResCode{
		Code: code,
		Data: data,
	}
	output, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, string(output))
}