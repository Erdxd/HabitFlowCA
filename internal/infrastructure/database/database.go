package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb(UrlDb string) (*sql.DB, error) {
	PsqlInfo := UrlDb
	var err error
	db, err = sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Println("Failed to connect to the database with your data")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to the database")
		return nil, err
	}
	return db, nil

}
