package middlewares

import (
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
	"github.com/boj/redistore"
)

var store *redistore.RediStore

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			store = database.GetSessionStore()
			session, err := store.Get(r, database.SessionID)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if session.Values["userID"] == nil {
				http.Error(w, "User not authenticated", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
}
