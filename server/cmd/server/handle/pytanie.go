package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"zawodowe-quiz/pkg/slices"
	"zawodowe-quiz/pkg/typy"
	"zawodowe-quiz/pkg/typy/kategorie"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Pytanie(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kat := czytajKategorie(r)
		query := fmt.Sprintf("SELECT Id,Pytanie,OdpA,OdpB,OdpC,OdpD, Obrazek,Poprawna FROM %s ORDER BY RANDOM() LIMIT 1;", kat)
		row := db.QueryRow(context.Background(), query)
		// z jakiegoś powodu scan działa tylko po podaniu każdego pola oddzielnie; nie można podać całego pytania (typy.Pytanie) jako celu skanu
		var p typy.Pytanie
		row.Scan(&p.Id, &p.Pytanie, &p.OdpA, &p.OdpB, &p.OdpC, &p.OdpD, &p.Obrazek, &p.Poprawna)
		resJson, _ := json.Marshal(&p)
		w.Write(resJson)
	}
}

func WszystkiePytania(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kat := r.URL.Query().Get("kategoria")
		if !slices.CzyZawiera(kategorie.WszystkieKategorie, kat) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		query := fmt.Sprintf("SELECT Id,Pytanie,OdpA,OdpB,OdpC,OdpD,Obrazek,Poprawna FROM %s ORDER BY Id;", kat)
		rows, err := db.Query(context.Background(), query)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		var pytania []typy.Pytanie
		for rows.Next() {
			var p typy.Pytanie
			rows.Scan(&p.Id, &p.Pytanie, &p.OdpA, &p.OdpB, &p.OdpC, &p.OdpD, &p.Obrazek, &p.Poprawna)
			pytania = append(pytania, p)
		}
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
