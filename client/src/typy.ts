export interface Pytanie {
	id: number;
	pytanie: string;
	odpA: string;
	odpB: string;
	odpC: string;
	odpD: string;
	obrazek: string;
	poprawna: number;
}

export type kategoria = "e12" | "e13" | "e14" | "ee08" | "ee09";
