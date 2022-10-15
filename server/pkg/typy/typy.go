package typy

type Pytanie struct {
	Id        int    `json:"id"`
	Kategoria string `json:"kategoria"`
	Pytanie   string `json:"pytanie"`
	OdpA      string `json:"odpA"`
	OdpB      string `json:"odpB"`
	OdpC      string `json:"odpC"`
	OdpD      string `json:"odpD"`
	Obrazek   string `json:"obrazek"`
	Poprawna  string `json:"poprawna"`
}
