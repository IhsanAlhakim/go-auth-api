package auth

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

func VerifyPassword(w http.ResponseWriter, hash string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			http.Error(w, "Incorrect sign in credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return err
	}
	return nil
}
