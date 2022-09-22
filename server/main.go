package main

import (
	"zawodowe-quiz/pkg/server"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("baza.db"), &gorm.Config{})
	if err != nil {
		panic("Nie udało się połączyć z bazą danych")
	}
	server.Init(db)
}
