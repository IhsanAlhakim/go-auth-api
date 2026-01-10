package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(w http.ResponseWriter, password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
