package routes

import (
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/handlers"
	"github.com/IhsanAlhakim/go-auth-api/pkg/mux"
)

func Register(mux *mux.Mux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Connection OK"))
	})
	mux.HandleFunc("GET /users/{id}", handlers.GetUser)
}
