import { defineStore } from "pinia";
import { Pytanie } from "../typy";

export const useModal = defineStore("modal-store", {
	state: () => ({
		pytanie: null as Pytanie | null,
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
