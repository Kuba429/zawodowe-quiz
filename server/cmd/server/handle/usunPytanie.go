package handle

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"zawodowe-quiz/cmd/zdjecie"
	"zawodowe-quiz/pkg/slices"
	"zawodowe-quiz/pkg/typy/kategorie"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UsunPytanie(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kat := r.URL.Query().Get("kategoria")
		if !slices.CzyZawiera(kategorie.WszystkieKategorie, kat) {
			http.Error(w, "Nie ma takiej kategorii", http.StatusBadRequest)
			return
		}
		idPytaniaRaw := r.URL.Query().Get("idPytania")
		idPytania, err := strconv.Atoi(idPytaniaRaw)
		if err != nil || len(idPytaniaRaw) < 1 || len(kat) < 1 {
			http.Error(w, "Nie można odczytać potrzebnych informacji o pytaniu", http.StatusInternalServerError)
		}
		zdjecie.Usun(db, idPytania, kat)
		query := fmt.Sprintf("DELETE FROM %s WHERE Id=%d;", kat, idPytania)
		if _, err := db.Exec(context.Background(), query); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(fmt.Sprintf("Usunięto pytanie o id %d i kategorii %s", idPytania, kat)))
	}
}
