package baza

import "gorm.io/gorm"

type Pytanie struct {
	Pytanie  string `json:"pytanie" gorm:"not null"`
	OdpA     string `json:"odpA"`
	OdpB     string `json:"odpB"`
	OdpC     string `json:"odpC"`
	OdpD     string `json:"odpD"`
	Obrazek  string `json:"obrazek"`
	Poprawna string `json:"poprawna" gorm:"not null"`
}

// wszystkie poniższe structy są identyczne
type E12 struct {
	gorm.Model
	Pytanie
}
type E13 struct {
	gorm.Model
	Pytanie
}
type E14 struct {
	gorm.Model
	Pytanie
}
type Ee08 struct {
	gorm.Model
	Pytanie
}
type Ee09 struct {
	gorm.Model
	Pytanie
}
