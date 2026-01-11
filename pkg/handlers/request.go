package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeRequestBody(w http.ResponseWriter, r *http.Request, payload any) error {
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		if err == io.EOF {
			http.Error(w, "Request body must not be empty", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return err
	}
	return nil
}
