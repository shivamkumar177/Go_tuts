package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConn() *sql.DB {
	config := ParseConfig()
	dbConnectionUrl := config.DbConnectionUrl
	db, err := sql.Open("postgres", dbConnectionUrl)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
