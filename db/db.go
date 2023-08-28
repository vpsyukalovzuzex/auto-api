package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Create() *sql.DB {
	var db *sql.DB
	var e error
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBNAME"))
	db, e = sql.Open("postgres", info)
	if e != nil {
		log.Fatalln(e)
	}
	e = db.Ping()
	if e != nil {
		log.Fatalln(e)
	}
	return db
}

func Close(db *sql.DB) {
	e := db.Close()
	if e != nil {
		log.Fatalln(e)
	}
}
