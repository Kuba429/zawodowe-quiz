package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"zawodowe-quiz/pkg/slices"
	"zawodowe-quiz/pkg/typy"
	"zawodowe-quiz/pkg/typy/kategorie"

	_ "github.com/mattn/go-sqlite3"
)

func Pytanie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kat := czytajKategorie(r)
		query := fmt.Sprintf("SELECT Id,Pytanie,OdpA,OdpB,OdpC,OdpD, Obrazek,Poprawna FROM %s ORDER BY RANDOM() LIMIT 1;", kat)
		rows := db.QueryRow(query)

		var p typy.Pytanie
		// z jakiegoś powodu scan działa tylko po podaniu każdego pola oddzielnie; nie można podać całego pytania (typy.Pytanie) jako celu skanu
		rows.Scan(&p.Id, &p.Pytanie, &p.OdpA, &p.OdpB, &p.OdpC, &p.OdpD, &p.Obrazek, &p.Poprawna)
		resJson, _ := json.Marshal(&p)
		w.Write(resJson)
	}
}

func WszystkiePytania(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kat := r.URL.Query().Get("kategoria")
		if !slices.CzyZawiera(kategorie.WszystkieKategorie, kat) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		query := fmt.Sprintf("SELECT Id,Pytanie,OdpA,OdpB,OdpC,OdpD,Obrazek,Poprawna FROM %s ORDER BY Id;", kat)
		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var pytania []typy.Pytanie
		for rows.Next() {
			var p typy.Pytanie
			rows.Scan(&p.Id, &p.Pytanie, &p.OdpA, &p.OdpB, &p.OdpC, &p.OdpD, &p.Obrazek, &p.Poprawna)
			pytania = append(pytania, p)
		}
		resJson, _ := json.Marshal(&pytania)
		w.Write(resJson)
	}
}

func czytajKategorie(r *http.Request) string {
	katQuery := r.URL.Query().Get("kategoria")
	queryOk := slices.CzyZawiera(kategorie.WszystkieKategorie, katQuery)
	if !queryOk {
		src := rand.NewSource(time.Now().UnixMicro())
		losowaLiczba := rand.New(src).Intn(len(kategorie.WszystkieKategorie))
		katQuery = kategorie.WszystkieKategorie[losowaLiczba]
	}
	return katQuery
}
