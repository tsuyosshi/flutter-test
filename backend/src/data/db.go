package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=postgres port=5432 user=user password=password dbname=flutter-test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func TestPost() (user_id int, user_name string) {
	statement := "insert into users (name) values ($1) returning id, name"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow("first_user").Scan(&user_id, &user_name)
	return
}
