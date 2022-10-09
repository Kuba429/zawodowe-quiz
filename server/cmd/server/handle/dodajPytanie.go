package handle

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"zawodowe-quiz/cmd/zdjecie"
	"zawodowe-quiz/pkg/typy"

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
		pytanie.Id, _ = strconv.Atoi(r.Form.Get("id"))
		pytanie.Poprawna, _ = strconv.Atoi(r.Form.Get("poprawna"))
		sciezka, err := zdjecie.Zapisz(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pytanie.Obrazek = sciezka

		fmt.Printf("%#v\n", pytanie)
		w.Write([]byte("aaa"))
	}
}
