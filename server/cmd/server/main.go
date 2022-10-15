package server

import (
	"log"
	"net/http"
	"zawodowe-quiz/cmd/server/handle"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(db *pgxpool.Pool) {
	fs := http.FileServer(http.Dir("./zdjecia"))
	http.Handle("/zdjecia/", http.StripPrefix("/zdjecia/", fs))

	http.HandleFunc("/pytanie", handle.Pytanie(db))
	http.HandleFunc("/wszystkie-pytania", handle.WszystkiePytania(db))
	http.HandleFunc("/baza/reset", handle.ResetBazy(db))
	http.HandleFunc("/update-pytanie", handle.UpdatePytanie(db))
	http.HandleFunc("/usun-pytanie", handle.UsunPytanie(db))
	http.HandleFunc("/dodaj-pytanie", handle.DodajPytanie(db))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
