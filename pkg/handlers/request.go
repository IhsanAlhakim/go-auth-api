package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeRequestBody(w http.ResponseWriter, r *http.Request, payload any) error {
	err := json.NewDecoder(r.Body).Decode(payload)

	if err != nil {
		if err == io.EOF {
			http.Error(w, "Request body must not be empty", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return err
	}

	switch {
	case err == io.EOF:

	case err != nil:

	}
	return nil
}
