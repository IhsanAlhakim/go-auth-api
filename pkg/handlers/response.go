package handlers

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

var (
	ServerError = "The server encountered an error, please try again later"
)

func Response(w http.ResponseWriter, payload Payload, statusCode int) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ServerError))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}
