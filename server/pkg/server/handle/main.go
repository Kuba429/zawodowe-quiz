package handle

import (
	"encoding/json"
	"net/http"
	"zawodowe-quiz/pkg/baza"

	"gorm.io/gorm"
)

func ResetBazy(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baza.Reset(db)
		w.Write([]byte("Baza zresetowana"))
	}
}

func Pytanie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res baza.E12
		db.First(&res)
		resJson, _ := json.Marshal(&res)
		w.Write(resJson)
	}
}
