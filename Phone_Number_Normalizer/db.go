package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConn() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://learn_db_owner:******@ep-muddy-sunset-a1lkdgwr.ap-southeast-1.aws.neon.tech/learn_db?sslmode=require")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
