package baza

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Pytania struct {
	E12  []E12  `json:"e12"`
	E13  []E13  `json:"e13"`
	E14  []E14  `json:"e14"`
	Ee08 []Ee08 `json:"ee08"`
	Ee09 []Ee09 `json:"ee09"`
}

func WczytajPytania() (Pytania, error) {
	plikJson, err := os.Open("pytaniaDomyslne.json")
	if err != nil {
		return Pytania{}, errors.New("nie można otworzyć pliku z pytaniami")
	}
	defer plikJson.Close()

	var pytania Pytania
	bajty, _ := ioutil.ReadAll(plikJson)
	json.Unmarshal(bajty, &pytania)

	return pytania, nil
}
