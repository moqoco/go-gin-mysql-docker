package model

import (
	"fmt"
	"log"
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

func Conn() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", "root", "root", "db", "3306", "test_db")
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Database connected.")

	return db
}