package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/internal/config"
	"github.com/IhsanAlhakim/go-auth-api/internal/database"
	"github.com/IhsanAlhakim/go-auth-api/internal/mux"
	"github.com/IhsanAlhakim/go-auth-api/internal/routes"
)

var db *sql.DB

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Panicf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	store, err := database.NewSessionStore(cfg)
	if err != nil {
		log.Panicf("Failed to create session store: %v", err)
	}
	defer store.Close()

	mux := mux.New()

	routes.Register(mux)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	log.Println("Server started at localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
