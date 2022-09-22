package handle

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"zawodowe-quiz/pkg/baza"

	"gorm.io/gorm"
)

func ResetBazy(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baza.Reset(db)
		w.Write([]byte("Baza zresetowana"))
	}
}

func Pytanie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kwal := czytajKwalifikacje(r)
		var pytanie baza.Pytanie
		query := fmt.Sprintf("SELECT * FROM %s ORDER BY RANDOM() LIMIT 1;", kwal)
		db.Raw(query).Scan(&pytanie)

		resJson, _ := json.Marshal(&pytanie)

		w.Write(resJson)
	}
}

func czytajKwalifikacje(r *http.Request) string {
	kwalifikacje := [5]string{"e12", "e13", "e14", "ee08", "ee09"}
	kwalQuery := r.URL.Query().Get("kwal")
	queryOk := false
	for _, k := range kwalifikacje {
		if kwalQuery == k {
			queryOk = true
			break
		}
	}
	if !queryOk {
		src := rand.NewSource(time.Now().UnixMicro())
		losowaLiczba := rand.New(src).Intn(len(kwalifikacje))
		kwalQuery = kwalifikacje[losowaLiczba]
	}
	return kwalQuery
}
