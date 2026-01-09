package handlers

import (
	"database/sql"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

var db *sql.DB

type P = Payload

func GetUser(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()
	id := r.PathValue("id")

	var user User

	row := db.QueryRow("SELECT username, email FROM users WHERE id = ?", id)
	if err := row.Scan(&user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			Response(w, P{Message: "Data not found"}, http.StatusNotFound)
		}
		Response(w, P{Message: ServerError}, http.StatusInternalServerError)
	}
	Response(w, P{Data: user}, http.StatusOK)
}
