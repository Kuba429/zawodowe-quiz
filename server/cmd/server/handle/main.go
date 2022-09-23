package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
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

func Pytanie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kwal := czytajKwalifikacje(r)
		query := fmt.Sprintf("SELECT Id,Pytanie,OdpA,OdpB,OdpC,OdpD, Obrazek,Poprawna FROM %s ORDER BY RANDOM() LIMIT 1;", kwal)
		rows := db.QueryRow(query)

		var p typy.Pytanie
		// z jakiegoś powodu scan działa tylko po podaniu każdego pola oddzielnie; nie można podać całego pytania (typy.Pytanie) jako celu skanu
		rows.Scan(&p.Id, &p.Pytanie, &p.OdpA, &p.OdpB, &p.OdpC, &p.OdpD, &p.Obrazek, &p.Poprawna)
		resJson, _ := json.Marshal(&p)
		w.Write(resJson)
	}
}

func czytajKwalifikacje(r *http.Request) kategorie.Kategoria {
	kwalQuery := kategorie.Kategoria(r.URL.Query().Get("kwal"))
	queryOk := false
	for _, k := range kategorie.WszystkieKategorie {
		if kwalQuery == k {
			queryOk = true
			break
		}
	}
	if !queryOk {
		src := rand.NewSource(time.Now().UnixMicro())
		losowaLiczba := rand.New(src).Intn(len(kategorie.WszystkieKategorie))
		kwalQuery = kategorie.WszystkieKategorie[losowaLiczba]
	}
	return kwalQuery
}
