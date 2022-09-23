package main

import (
	"database/sql"
	"zawodowe-quiz/cmd/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./baza.db")
	if err != nil {
		panic("Nie udało się połączyć z bazą danych")
	}
	server.Init(db)

}
