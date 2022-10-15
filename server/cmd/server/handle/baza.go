package handle

import (
	"net/http"
	"zawodowe-quiz/pkg/baza"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ResetBazy(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		baza.Reset(db)
		w.Write([]byte("Baza zresetowana"))
	}
}
