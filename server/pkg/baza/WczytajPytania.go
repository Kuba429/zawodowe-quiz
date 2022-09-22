package baza

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Pytanie struct {
	Pytanie  string `json:"pytanie"`
	OdpA     string `json:"odpA"`
	OdpB     string `json:"odpB"`
	OdpC     string `json:"odpC"`
	OdpD     string `json:"odpD"`
	Obrazek  string `json:"obrazek"`
	Poprawna string `json:"poprawna"`
}
type Pytania struct {
	E12 []E12 `json:"e12"`
	E13 []E13 `json:"e13"`
	E14 []E14 `json:"e14"`
	Ee8 []Ee8 `json:"ee8"`
	Ee9 []Ee9 `json:"ee9"`
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
