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

func WszystkiePytania(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kwal := kategorie.Kategoria(r.URL.Query().Get("kwal"))
		if !slices.CzyZawiera(kategorie.WszystkieKategorie, kwal) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		query := fmt.Sprintf("SELECT Id,Pytanie,OdpA,OdpB,OdpC,OdpD,Obrazek,Poprawna FROM %s ORDER BY Id;", kwal)
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

func czytajKwalifikacje(r *http.Request) kategorie.Kategoria {
	kwalQuery := kategorie.Kategoria(r.URL.Query().Get("kwal"))
	queryOk := slices.CzyZawiera(kategorie.WszystkieKategorie, kwalQuery)
	if !queryOk {
		src := rand.NewSource(time.Now().UnixMicro())
		losowaLiczba := rand.New(src).Intn(len(kategorie.WszystkieKategorie))
		kwalQuery = kategorie.WszystkieKategorie[losowaLiczba]
	}
	return kwalQuery
}
