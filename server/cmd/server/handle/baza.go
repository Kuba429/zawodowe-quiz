package handle

import (
	"database/sql"
	"net/http"
	"zawodowe-quiz/pkg/baza"
	"zawodowe-quiz/pkg/zdjecie"

	_ "github.com/mattn/go-sqlite3"
)

func ResetBazy(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		baza.Reset(db)
		w.Write([]byte("Baza zresetowana"))
	}
}

func UpdatePytanie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.ParseMultipartForm(1024 * 1024)
		nowePytanie := r.Form
		if !nowePytanie.Has("obrazek") {
			sciezka, err := zdjecie.Zapisz(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			nowePytanie.Set("obrazek", sciezka)
		}

		w.Write([]byte("aaa"))
	}
}
