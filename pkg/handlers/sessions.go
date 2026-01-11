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

	DecodeRequestBody(w, r, &credentials)

	var user User

	row := db.QueryRow("SELECT username, password FROM users WHERE email = ?", credentials.Email)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			Response(w, P{Message: "User not found"}, http.StatusNotFound)
		}
		Response(w, P{Message: ServerError}, http.StatusInternalServerError)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	switch {
	case err == bcrypt.ErrMismatchedHashAndPassword:
		Response(w, P{Message: "Incorrect Sign In Credentials"}, http.StatusUnauthorized)
	case err != nil:
		Response(w, P{Message: ServerError}, http.StatusInternalServerError)

	}

	session, err := store.Get(r, database.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.Values["userID"] = user.Username
	session.Save(r, w)

	Response(w, P{Message: "Sign In Successfull"}, http.StatusOK)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, database.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
	Response(w, P{Message: "Sign Out Successful"}, http.StatusOK)
}
