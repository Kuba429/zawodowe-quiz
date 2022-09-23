package server

import (
	"database/sql"
	"log"
	"net/http"
	"zawodowe-quiz/cmd/server/handle"
)

func Init(db *sql.DB) {
	http.HandleFunc("/pytanie", handle.Pytanie(db))
	http.HandleFunc("/baza/reset", handle.ResetBazy(db))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
