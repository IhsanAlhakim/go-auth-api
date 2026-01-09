package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/mux"
)

func main() {
	mux := mux.New()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Connection OK"))
	})

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	fmt.Println("Server started at localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Println("Shutting down server...")
	}
}
