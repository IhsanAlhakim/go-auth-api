package handlers

import (
	"database/sql"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
	"github.com/IhsanAlhakim/go-auth-api/pkg/utils"
	"github.com/boj/redistore"
)

var store *redistore.RediStore

func SignIn(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()
	store = database.GetSessionStore()

	var credentials = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := DecodeRequestBody(w, r, &credentials); err != nil {
		return
	}

	if err := utils.CheckStructEmptyProperty(credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.CheckStructWhitespaceProperty(credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	if err := utils.IsPasswordCorrect(w, user.Password, credentials.Password); err != nil {
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
