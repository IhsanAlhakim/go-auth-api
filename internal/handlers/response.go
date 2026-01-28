package handlers

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type P = Payload

func Response(w http.ResponseWriter, payload Payload, statusCode int) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}
