package handlers

import (
	"database/sql"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
	"github.com/IhsanAlhakim/go-auth-api/pkg/utils"
)

type User struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var db *sql.DB

func GetUser(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()
	userId := r.PathValue("id")

	var user User

	row := db.QueryRow("SELECT username, email FROM users WHERE id = ?", userId)
	if err := row.Scan(&user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	Response(w, P{Data: user}, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()

	var user User

	if err := DecodeRequestBody(w, r, &user); err != nil {
		return
	}

	if err := utils.CheckStructEmptyProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.CheckStructWhitespaceProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.GenerateHashPassword(w, user.Password)
	if err != nil {
		return
	}

	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Response(w, P{Message: "User Created!"}, http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()
	userId := r.PathValue("id")

	var user = struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}{}

	if err := DecodeRequestBody(w, r, &user); err != nil {
		return
	}

	if err := utils.CheckStructEmptyProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.CheckStructWhitespaceProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", user.Username, user.Email, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Response(w, P{Message: "User Updated!"}, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()
	userId := r.PathValue("id")

	var user User

	row := db.QueryRow("SELECT username FROM users WHERE id = ?", userId)
	if err := row.Scan(&user.Username); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	_, err := db.Exec("DELETE FROM users WHERE id = ?", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Response(w, P{Message: "User Deleted!"}, http.StatusOK)
}
