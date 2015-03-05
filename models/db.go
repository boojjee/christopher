package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

// Creates and tests database connection
func ConnectDb() {
	var err error
	ds := "root:AdminAdmin@/christopher?parseTime=true"
	DB, err = sql.Open("mysql", ds)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()
}

// Closes database connection
func CloseDb() {
	DB.Close()
}
