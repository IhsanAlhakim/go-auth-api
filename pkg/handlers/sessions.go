package handlers

import (
	"database/sql"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
	"github.com/boj/redistore"
	"golang.org/x/crypto/bcrypt"
)

var store *redistore.RediStore

func SignIn(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()
	store = database.GetSessionStore()

	var credentials User

	if err := DecodeRequestBody(w, r, &credentials); err != nil {
		return
	}

	var user User

	row := db.QueryRow("SELECT username, password FROM users WHERE email = ?", credentials.Email)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	switch {
	case err == bcrypt.ErrMismatchedHashAndPassword:
		http.Error(w, "Incorrect sign in credentials", http.StatusUnauthorized)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := store.Get(r, database.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["userID"] = user.Username
	session.Save(r, w)

	Response(w, P{Message: "Sign In Successfull"}, http.StatusOK)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, database.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
	Response(w, P{Message: "Sign Out Successful"}, http.StatusOK)
}
