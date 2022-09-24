import { defineStore } from "pinia";
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
			const url = new URL("http://localhost:3000/pytanie");
			this.kategoria &&
				(url.search = new URLSearchParams({
					kwal: this.kategoria,
				}).toString());

			try {
				const res = await fetch(url);
				const text = await (await res.blob()).text();
				const p: Pytanie = JSON.parse(text);
				this.pytanie = p;
				this.odpowiedzi = [p.odpA, p.odpB, p.odpC, p.odpD];
				this.poprawnaOdp = this.odpowiedzi[p.poprawna];
				this.wybranaOdp = "";
				this.status = "success";
			} catch (error) {
				console.error(error);
				this.status = "error";
			}
		},
	},
});
