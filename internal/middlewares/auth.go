package middlewares

import (
	"net/http"

	"github.com/boj/redistore"
)

var store *redistore.RediStore

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session, err := m.store.Get(r, m.cfg.SessionID)

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
