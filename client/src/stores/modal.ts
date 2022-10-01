import { defineStore } from "pinia";
import { kategoria, Pytanie } from "../typy";

export const useModal = defineStore("modal-store", {
	state: () => ({
		pytanie: null as Pytanie | null,
		kategoria: "e12" as kategoria,
	}),
	actions: {
		zamknij() {
			this.pytanie = null;
		},
		ustaw(p: Pytanie) {
			this.pytanie = p;
		},
	},
});
