package config

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	Once sync.Once
)

func PgConnect() *sql.DB {
	//
	if DB != nil {
		return DB
	}

	Once.Do(func() {
		db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		if db.Ping() != nil {
			log.Fatal(errors.New("CANNOT PING DATABASE"))
		}

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		DB = db
	})

	return DB

}
