package handle

import (
	"database/sql"
	"fmt"
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
		query := fmt.Sprintf(`
		UPDATE %s SET
		Pytanie = "%s",
		OdpA = "%s",
		OdpB = "%s",
		OdpC = "%s",
		OdpD = "%s",
		Obrazek = "%s",
		Poprawna = %s
		WHERE Id = %s;
		`, nowePytanie["kategoria"][0], nowePytanie["pytanie"][0], nowePytanie["odpA"][0], nowePytanie["odpB"][0], nowePytanie["odpC"][0], nowePytanie["odpD"][0], nowePytanie["obrazek"][0], nowePytanie["poprawna"][0], nowePytanie["id"][0])

		if _, err := db.Exec(query); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte("aaa"))
	}
}
