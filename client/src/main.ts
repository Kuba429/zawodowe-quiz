import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { createRouter, createWebHistory } from "vue-router";
import Home from "./routes/Home.vue";
import Pytanie from "./routes/Pytanie.vue";

const router = createRouter({
	history: createWebHistory(),
	routes: [
		{ path: "/", name: "Home", component: Home },
		{ path: "/pytanie/:kategoria?", name: "Pytanie", component: Pytanie },
	],
});

createApp(App).use(router).mount("#app");
