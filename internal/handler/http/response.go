package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func JSON(w http.ResponseWriter, message string, data interface{}) {
	response := Response{
		Data:    data,
		Message: message,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
