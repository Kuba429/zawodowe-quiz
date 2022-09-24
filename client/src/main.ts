import { createRouter, createWebHistory } from "vue-router";
import { createApp } from "vue";
import { createPinia } from "pinia";
import "./style.css";
import App from "./App.vue";
import StronaGlowna from "./routes/StronaGlowna.vue";
import Kategoria from "./routes/Kategoria.vue";

const router = createRouter({
	history: createWebHistory(),
	routes: [
		{ path: "/", name: "Home", component: StronaGlowna },
		{ path: "/pytanie/:kategoria?", name: "Pytanie", component: Kategoria },
	],
});

createApp(App).use(router).use(createPinia()).mount("#app");
