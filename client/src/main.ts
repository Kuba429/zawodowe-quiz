import { createRouter, createWebHistory } from "vue-router";
import { createApp } from "vue";
import { createPinia } from "pinia";
import "./style/global.scss";
import App from "./App.vue";
import StronaGlowna from "./routes/StronaGlowna.vue";
import Kategoria from "./routes/Kategoria.vue";
import Panel from "./routes/Panel.vue";
import Dodaj from "./routes/Dodaj.vue";

const router = createRouter({
	history: createWebHistory(),
	routes: [
		{ path: "/", name: "Home", component: StronaGlowna },
		{ path: "/pytanie/:kategoria?", name: "Pytanie", component: Kategoria },
		{ path: "/panel", name: "Panel", component: Panel },
		{ path: "/dodaj", name: "Dodaj", component: Dodaj },
	],
});

createApp(App).use(router).use(createPinia()).mount("#app");

export const HOST = import.meta.env.VITE_HOST || "http://localhost:3000";
export const API = import.meta.env.VITE_API || "http://localhost:3000";
