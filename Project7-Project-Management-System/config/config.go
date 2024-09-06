package config

import (
	"database/sql"
	"log"
	_ "os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dsn := "user:password@tcp(localhost:3306)/dbname"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	log.Println("Database connected successfully")
}

func GetDB() *sql.DB {
	return DB
}
