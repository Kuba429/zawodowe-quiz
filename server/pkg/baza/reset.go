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
	db.Migrator().DropTable(&E12{})
	db.Migrator().DropTable(&E13{})
	db.Migrator().DropTable(&E14{})
	db.Migrator().DropTable(&Ee8{})
	db.Migrator().DropTable(&Ee9{})
	// migruj modele z pliku baza/tabele
	db.AutoMigrate(&E12{})
	db.AutoMigrate(&E13{})
	db.AutoMigrate(&E14{})
	db.AutoMigrate(&Ee8{})
	db.AutoMigrate(&Ee9{})
	// umieść domyślne pytania w nowo utworzonych tabelach (zostały przed chwilą usunięte i utworzone ponownie)
	db.CreateInBatches(&pytania.E12, len(pytania.E12))
	db.CreateInBatches(&pytania.E13, len(pytania.E13))
	db.CreateInBatches(&pytania.E14, len(pytania.E14))
	db.CreateInBatches(&pytania.Ee8, len(pytania.Ee8))
	db.CreateInBatches(&pytania.Ee9, len(pytania.Ee9))

	log.Println("Baza pytań została zresetowana")
}
