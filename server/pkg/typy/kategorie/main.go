package kategorie

type Kategoria string

const (
	E12  Kategoria = "e12"
	E13  Kategoria = "e13"
	E14  Kategoria = "e14"
	Ee08 Kategoria = "ee08"
	Ee09 Kategoria = "ee09"
)

// brak statycznego rozmiaru aby można było użyć w funkcji "czyZawiera" (cmd/server/handle/main)
var WszystkieKategorie = []Kategoria{E12, E13, E14, Ee08, Ee09}
