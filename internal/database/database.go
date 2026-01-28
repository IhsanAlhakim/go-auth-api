package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.User = os.Getenv("DB_USERNAME")
	config.Passwd = os.Getenv("DB_PASSWORD")
	config.Net = "tcp"
	config.Addr = "127.0.0.1:3306"
	config.DBName = os.Getenv("DB_NAME")

	var err error

	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}
