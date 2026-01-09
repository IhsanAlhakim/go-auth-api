package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/pkg/database"
	"github.com/IhsanAlhakim/go-auth-api/pkg/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	mux := mux.New()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Connection OK"))
	})

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	log.Println("Server started at localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
		log.Fatal("Shutting down server...")
	}
}
