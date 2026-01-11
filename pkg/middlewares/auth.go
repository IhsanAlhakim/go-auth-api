package middlewares

import (
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
	"github.com/IhsanAlhakim/go-auth-api/pkg/handlers"
	"github.com/boj/redistore"
)

var store *redistore.RediStore

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			store = database.GetSessionStore()
			session, _ := store.Get(r, database.SessionID)

			if session.Values["userID"] == nil {
				handlers.Response(w, handlers.Payload{Message: "User not authenticated"}, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
}
