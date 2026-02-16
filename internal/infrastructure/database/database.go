package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb(UrlDb string) (*sql.DB, error) {
	PsqlInfo := UrlDb
	var err error
	db, err = sql.Open("postgres", PsqlInfo)
	if err != nil {

		return nil, err
	}
	err = db.Ping()
	if err != nil {

		return nil, err
	}
	return db, nil

}
