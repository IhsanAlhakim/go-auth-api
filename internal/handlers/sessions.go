package handlers

import (
	"database/sql"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/internal/auth"
	"github.com/IhsanAlhakim/go-auth-api/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	var credentials = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := BindJSON(r, &credentials); err != nil {
		if err == ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := validation.CheckStructEmptyProperty(credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validation.CheckStructWhitespaceProperty(credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User

	row := h.db.QueryRow("SELECT username, password FROM users WHERE email = ?", credentials.Email)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := auth.VerifyPassword(user.Password, credentials.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	session, err := h.store.Get(r, h.cfg.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["userID"] = user.Username
	session.Save(r, w)

	Response(w, P{Message: "Sign In Successfull"}, http.StatusOK)
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	session, err := h.store.Get(r, h.cfg.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
	Response(w, P{Message: "Sign Out Successful"}, http.StatusOK)
}
