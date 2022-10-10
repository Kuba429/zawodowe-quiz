package handle

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"zawodowe-quiz/cmd/zdjecie"
	"zawodowe-quiz/pkg/slices"
	"zawodowe-quiz/pkg/typy"
	"zawodowe-quiz/pkg/typy/kategorie"

	_ "github.com/mattn/go-sqlite3"
)

func DodajPytanie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.ParseMultipartForm(1024 * 1024)

		pytanie := typy.Pytanie{
			Pytanie:   r.Form.Get("pytanie"),
			Kategoria: r.Form.Get("kategoria"),
			OdpA:      r.Form.Get("odpA"),
			OdpB:      r.Form.Get("odpB"),
			OdpC:      r.Form.Get("odpC"),
			OdpD:      r.Form.Get("odpD"),
			Obrazek:   r.Form.Get("obrazek"),
		}
		if !slices.CzyZawiera(kategorie.WszystkieKategorie, kategorie.Kategoria(pytanie.Kategoria)) {
			http.Error(w, "Nie ma takiej kategorii", http.StatusBadRequest)
			return
		}
		pytanie.Poprawna, _ = strconv.Atoi(r.Form.Get("poprawna"))
		sciezka, err := zdjecie.Zapisz(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pytanie.Obrazek = sciezka

		query := fmt.Sprintf("INSERT INTO %s (Pytanie, Kategoria, OdpA, OdpB, OdpC, OdpD, Obrazek, Poprawna) VALUES ('%s','%s','%s','%s','%s','%s','%s',%d)", pytanie.Kategoria, pytanie.Pytanie, pytanie.Kategoria, pytanie.OdpA, pytanie.OdpB, pytanie.OdpC, pytanie.OdpD, pytanie.Obrazek, pytanie.Poprawna)
		if _, err := db.Exec(query); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte("Dodano pytanie"))
	}
}
