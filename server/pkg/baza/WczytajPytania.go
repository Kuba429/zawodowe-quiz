package baza

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"zawodowe-quiz/pkg/typy"
	"zawodowe-quiz/pkg/typy/kategorie"
)

func WczytajPytania() (map[kategorie.Kategoria][]typy.Pytanie, error) {
	var pytania map[kategorie.Kategoria][]typy.Pytanie
	plikJson, err := os.Open("pytaniaDomyslne.json")
	if err != nil {
		return pytania, errors.New("nie można otworzyć pliku z pytaniami")
	}
	defer plikJson.Close()

	bajty, _ := ioutil.ReadAll(plikJson)
	json.Unmarshal(bajty, &pytania)

	return pytania, nil
}
