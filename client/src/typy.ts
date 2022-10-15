export interface Pytanie {
	id: number;
	kategoria?: kategoria;
	pytanie: string;
	odpA: string;
	odpB: string;
	odpC: string;
	odpD: string;
	obrazek: string;
	poprawna: string;
}

export type kategoria = "e12" | "e13" | "e14" | "ee08" | "ee09";

export const wszyskieKategorie: kategoria[] = [
	"e12",
	"e13",
	"e14",
	"ee08",
	"ee09",
];

export type statusType = "sukces" | "blad" | "ladowanie";
