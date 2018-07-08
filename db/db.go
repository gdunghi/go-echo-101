package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() *sql.DB {

	db, err := sql.Open("mysql", "go101:go10199@tcp(188.166.212.73:3306)/go_101")

	if err != nil {
		log.Println(err)
	}

	return db
}
