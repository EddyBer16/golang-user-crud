package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "user:pasword@tcp(host:port)/name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB CONNECTED")
	con = db
	return db
}
