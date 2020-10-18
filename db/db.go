package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// var db *sqlx.DB

type DB struct {
	*sqlx.DB
}

func DBConnect() (*DB, error) {
	// these properties can be read from config
	conString := "host=" + "localhost" + " user=" + "postgres" + " password=" + "postgres" + " dbname=" + "employee" + " sslmode=" + "disable"
	db, err := sqlx.Open("postgres", conString)

	if err != nil {
		log.Fatalln(err)
	}
	return &DB{db}, err
}

func SqlxConnect() (*DB, error) {
	// return db
	db, err := DBConnect()

	return db, err
}
