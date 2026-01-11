package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(w http.ResponseWriter, password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return hash, nil
}
