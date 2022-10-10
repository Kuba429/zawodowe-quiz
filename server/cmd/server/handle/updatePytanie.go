package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"zawodowe-quiz/cmd/zdjecie"
	"zawodowe-quiz/pkg/slices"
	"zawodowe-quiz/pkg/typy"
	"zawodowe-quiz/pkg/typy/kategorie"

	_ "github.com/mattn/go-sqlite3"
)

func UpdatePytanie(db *sql.DB) http.HandlerFunc {
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
		if !slices.CzyZawiera(kategorie.WszystkieKategorie, pytanie.Kategoria) {
			http.Error(w, "Nie ma takiej kategorii", http.StatusBadRequest)
			return
		}
		pytanie.Id, _ = strconv.Atoi(r.Form.Get("id"))
		pytanie.Poprawna, _ = strconv.Atoi(r.Form.Get("poprawna"))

		if !r.Form.Has("obrazek") {
			sciezka, err := zdjecie.Zapisz(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			pytanie.Obrazek = sciezka
		}
		// usun stare zdjecie zanim zostanie nadpisane
		zdjecie.Usun(db, pytanie.Id, pytanie.Kategoria)

		query := fmt.Sprintf(`
		UPDATE %s SET
		Pytanie = "%s",
		OdpA = "%s",
		OdpB = "%s",
		OdpC = "%s",
		OdpD = "%s",
		Obrazek = "%s",
		Poprawna = %d
		WHERE Id = %d;
		`, pytanie.Kategoria, pytanie.Pytanie, pytanie.OdpA, pytanie.OdpB, pytanie.OdpC, pytanie.OdpD, pytanie.Obrazek, pytanie.Poprawna, pytanie.Id)

		if _, err := db.Exec(query); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		pytanieJson, _ := json.Marshal(pytanie)
		w.Write(pytanieJson)
	}
}
