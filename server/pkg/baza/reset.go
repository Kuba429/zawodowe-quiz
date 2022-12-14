package baza

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"zawodowe-quiz/cmd/zdjecie"

	_ "github.com/mattn/go-sqlite3"
)

func Reset(db *sql.DB) {
	// otwórz plik z domyślnymi pytaniami. Jeśli się nie uda, zatrzymaj funckję tutaj zanim rozpocznie się proces usuwania tabel
	pytania, err := WczytajPytania()
	if err != nil {
		log.Println(err)
		return
	}
	// usuń wszystkie tabele
	for kat := range pytania {
		db.Exec(fmt.Sprintf("DROP TABLE %s;", kat))
	}
	// usun zdjęcia dodane przez użytkownikow (domyślne zdjęcia do pytań które nie zostaną usunięte są w folderze wyżej więc nie zostaną usunięte)
	if err := os.RemoveAll(zdjecie.Folder); err != nil {
		log.Println(err)
	}
	// utwórz tabele uzupełnij je pytaniami
	// połącz wszystkie kwerendy w jedną aby ograniczyć zapytania do bazy danych
	fullQuery := ""
	for key, val := range pytania {
		fullQuery += fmt.Sprintf("CREATE TABLE %s (%s);", key, PolaTabeli)
		for _, pytanie := range val {
			fullQuery += fmt.Sprintf(`
			INSERT INTO %s (Pytanie, Kategoria, OdpA, OdpB, OdpC, OdpD, Obrazek, Poprawna) VALUES ("%s","%s","%s","%s","%s","%s","%s","%s");`,
				key, pytanie.Pytanie, key, pytanie.OdpA, pytanie.OdpB, pytanie.OdpC, pytanie.OdpD, pytanie.Obrazek, pytanie.Poprawna)
		}
	}
	{
		_, err := db.Exec(fullQuery)
		if err != nil {
			log.Println(err)
			return
		}
	}
	log.Println("Baza pytań została zresetowana")
}

const PolaTabeli = `
Id INTEGER PRIMARY KEY AUTOINCREMENT,
Pytanie text NOT NULL,
Kategoria text,
OdpA text NOT NULL,
OdpB text NOT NULL,
OdpC text NOT NULL,
OdpD text NOT NULL,
Obrazek text,
Poprawna text NOT NULL
`
