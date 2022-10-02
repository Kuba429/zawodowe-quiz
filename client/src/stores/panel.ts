import { defineStore } from "pinia";
import { kategoria, Pytanie } from "../typy";

export const usePanel = defineStore("modal-store", {
	state: () => ({
		pytania: [] as Pytanie[],
		modalPytanie: null as Pytanie | null,
		kategoria: "e12" as kategoria,
	}),
	actions: {
		update(p: Pytanie) {
			const idx = this.pytania.findIndex((x) => x.id === p.id);
			this.pytania[idx] = p;
		},
		zamknij() {
			this.modalPytanie = null;
		},
		ustaw(p: Pytanie) {
			this.modalPytanie = p;
		},
		async fetchPytania() {
			const url = new URL("http://localhost:3000/wszystkie-pytania");
			url.search = new URLSearchParams({
				kwal: this.kategoria,
			}).toString();
			try {
				const res = await fetch(url);
				const text = await (await res.blob()).text();
				this.pytania = JSON.parse(text);
			} catch (error) {
				console.error(error);
			}
		},
	},
});
