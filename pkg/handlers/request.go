package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeRequestBody(w http.ResponseWriter, r *http.Request, payload any) {
	err := json.NewDecoder(r.Body).Decode(payload)

	switch {
	case err == io.EOF:
		Response(w, P{Message: "Request body must not be empty"}, http.StatusBadRequest)
	case err != nil:
		Response(w, P{Message: ServerError}, http.StatusInternalServerError)
	}
}
