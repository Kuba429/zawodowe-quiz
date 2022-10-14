import { defineStore } from "pinia";
import { API, HOST } from "../main";
import { Pytanie } from "../typy";

type status = "success" | "error" | "loading";

export const usePStore = defineStore("pytanie-store", {
	state: () => ({
		pytanie: null as Pytanie | null,
		kategoria: "",
		odpowiedzi: [] as string[],
		poprawnaOdp: "",
		wybranaOdp: "",
		status: "success" as status,
	}),
	actions: {
		async nowePytanie() {
			this.status = "loading";
			const url = new URL(`${API}/pytanie`);
			this.kategoria &&
				(url.search = new URLSearchParams({
					kategoria: this.kategoria,
				}).toString());

			try {
				const res = await fetch(url);
				const p: Pytanie = await res.json();
				if (p.obrazek) {
					p.obrazek = HOST + p.obrazek;
				}
				this.pytanie = p;
				let odpowiedzi = [p.odpA, p.odpB, p.odpC, p.odpD];

				this.poprawnaOdp = odpowiedzi[p.poprawna];
				// nie mieszaj jeśli odpowiedzi to litery A,B,C,D
				if (odpowiedzi[0] !== "A") {
					odpowiedzi = odpowiedzi
						.map((v) => ({
							v,
							idx: Math.random(),
						}))
						.sort((a, b) => b.idx - a.idx)
						.map(({ v }) => v);
				}
				this.odpowiedzi = odpowiedzi;
				this.wybranaOdp = "";
				this.status = "success";
			} catch (error) {
				console.error(error);
				this.status = "error";
			}
		},
	},
});
