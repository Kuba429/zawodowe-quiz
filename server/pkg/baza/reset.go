package baza

import (
	"context"
	"fmt"
	"log"
	"os"
	"zawodowe-quiz/cmd/zdjecie"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Reset(db *pgxpool.Pool) {
	// otwórz plik z domyślnymi pytaniami. Jeśli się nie uda, zatrzymaj funckję tutaj zanim rozpocznie się proces usuwania tabel
	pytania, err := WczytajPytania()
	if err != nil {
		log.Println(err)
		return
	}
	// usuń wszystkie tabele
	for kat := range pytania {
		db.Exec(context.Background(), fmt.Sprintf("DROP TABLE %s;", kat))
		db.Exec(context.Background(), fmt.Sprintf("CREATE TABLE %s (%s);", kat, PolaTabeli))
	}
	// usun zdjęcia dodane przez użytkownikow (domyślne zdjęcia do pytań które nie zostaną usunięte są w folderze wyżej więc nie zostaną usunięte)
	if err := os.RemoveAll(zdjecie.Folder); err != nil {
		log.Println(err)
	}
	// utwórz tabele uzupełnij je pytaniami
	// połącz wszystkie kwerendy w jedną aby ograniczyć zapytania do bazy danych
	fullQuery := ""
	for key, val := range pytania {
		for _, pytanie := range val {
			fullQuery += fmt.Sprintf(`
			INSERT INTO %s (Pytanie, Kategoria, OdpA, OdpB, OdpC, OdpD, Obrazek, Poprawna) VALUES ('%s','%s','%s','%s','%s','%s','%s','%d');`,
				key, pytanie.Pytanie, key, pytanie.OdpA, pytanie.OdpB, pytanie.OdpC, pytanie.OdpD, pytanie.Obrazek, pytanie.Poprawna)
		}
	}
	{
		_, err := db.Exec(context.Background(), fullQuery)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	log.Println("Baza pytań została zresetowana")
}

const PolaTabeli = `
Id SERIAL PRIMARY KEY,
Pytanie text NOT NULL,
Kategoria text,
OdpA text NOT NULL,
OdpB text NOT NULL,
OdpC text NOT NULL,
OdpD text NOT NULL,
Obrazek text,
Poprawna text NOT NULL
`
