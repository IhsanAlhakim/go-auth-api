package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
	"github.com/IhsanAlhakim/go-auth-api/pkg/mux"
	"github.com/IhsanAlhakim/go-auth-api/pkg/routes"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to loading .env file: %v", err)
	}

	db, err = database.Connect()
	if err != nil {
		log.Panicf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	store, err := database.NewSessionStore()
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
