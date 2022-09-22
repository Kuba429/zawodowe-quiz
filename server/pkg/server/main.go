package server

import (
	"log"
	"net/http"
	"zawodowe-quiz/pkg/server/handle"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	http.HandleFunc("/pytanie", handle.Pytanie(db))
	http.HandleFunc("/baza/reset", handle.ResetBazy(db))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
