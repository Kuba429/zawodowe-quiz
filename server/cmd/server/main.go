package server

import (
	"database/sql"
	"log"
	"net/http"
	"zawodowe-quiz/cmd/server/handle"
)

func Init(db *sql.DB) {
	fs := http.FileServer(http.Dir("./zdjecia"))
	http.Handle("/zdjecia/", http.StripPrefix("/zdjecia/", fs))

	http.HandleFunc("/pytanie", handle.Pytanie(db))
	http.HandleFunc("/wszystkie-pytania", handle.WszystkiePytania(db))
	http.HandleFunc("/baza/reset", handle.ResetBazy(db))
	http.HandleFunc("/update-pytanie", handle.UpdatePytanie(db))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
