package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDb() *sql.DB {
	connection := "user=meu_usuario dbname=meu_database password=meu_password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
