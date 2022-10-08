package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"zawodowe-quiz/cmd/zdjecie"
	"zawodowe-quiz/pkg/baza"
	"zawodowe-quiz/pkg/typy"
	"zawodowe-quiz/pkg/typy/kategorie"

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

		pytanie := typy.Pytanie{
			Pytanie:   r.Form.Get("pytanie"),
			Kategoria: r.Form.Get("kategoria"),
			OdpA:      r.Form.Get("odpA"),
			OdpB:      r.Form.Get("odpB"),
			OdpC:      r.Form.Get("odpC"),
			OdpD:      r.Form.Get("odpD"),
			Obrazek:   r.Form.Get("obrazek"),
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
		zdjecie.Usun(db, pytanie.Id, kategorie.Kategoria(pytanie.Kategoria))

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

func UsunPytanie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kwal := kategorie.Kategoria(r.URL.Query().Get("kwal"))
		idPytaniaRaw := r.URL.Query().Get("idPytania")
		idPytania, err := strconv.Atoi(idPytaniaRaw)
		if err != nil || len(idPytaniaRaw) < 1 || len(kwal) < 1 {
			http.Error(w, "Nie można odczytać potrzebnych informacji o pytaniu", http.StatusInternalServerError)
		}
		zdjecie.Usun(db, idPytania, kwal)
		query := fmt.Sprintf("DELETE FROM %s WHERE Id=%d;", kwal, idPytania)
		if _, err := db.Exec(query); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(fmt.Sprintf("Usunięto pytanie o id %d i kategorii %s", idPytania, kwal)))
	}
}
