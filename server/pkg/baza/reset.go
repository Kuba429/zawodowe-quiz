package baza

import (
	"log"

	"gorm.io/gorm"
)

func Reset(db *gorm.DB) {
	// otwórz plik z domyślnymi pytaniami. Jeśli się nie uda, zatrzymaj funckję tutaj zanim rozpocznie się proces usuwania tabel
	pytania, err := WczytajPytania()
	if err != nil {
		log.Println(err)
		return
	}
	// usuń wszystkie tabele
	db.Migrator().DropTable(&E12{}, &E13{}, &E14{}, &Ee08{}, &Ee09{})
	// migruj modele z pliku baza/tabele
	db.AutoMigrate(&E12{}, &E13{}, &E14{}, &Ee08{}, &Ee09{})
	// umieść domyślne pytania w nowo utworzonych tabelach (zostały przed chwilą usunięte i utworzone ponownie)
	db.CreateInBatches(&pytania.E12, len(pytania.E12))
	db.CreateInBatches(&pytania.E13, len(pytania.E13))
	db.CreateInBatches(&pytania.E14, len(pytania.E14))
	db.CreateInBatches(&pytania.Ee08, len(pytania.Ee08))
	db.CreateInBatches(&pytania.Ee09, len(pytania.Ee09))

	log.Println("Baza pytań została zresetowana")
}
