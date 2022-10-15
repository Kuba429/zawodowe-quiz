<script lang="ts" setup>
import { onMounted } from "vue";
import { useRoute } from "vue-router";
import Pytanie from "../components/Pytanie.vue";
import { usePStore } from "../stores/pytanie";

const store = usePStore();
const { kategoria } = useRoute().params;
onMounted(() => {
	store.$patch({ kategoria: kategoria as string });
	store.nowePytanie();
});
</script>
<template>
	<div class="grid">
		<h2>{{ kategoria || "Wszystkie kategorie" }}</h2>
		<button class="nowe-pytanie" @click="store.nowePytanie">
			Nowe pytanie
		</button>
	</div>
	<Pytanie v-if="store.pytanie && store.status === 'success'" />
</template>
<style lang="scss">
@use "../style/zmienne.scss" as *;
h2 {
	text-align: center;
	text-transform: capitalize;
}
.nowe-pytanie {
	display: block;
	margin: auto;
	margin-top: 10px;
	background-color: $niebieski;
	border-radius: 6px;
	border: none;
	color: white;
	padding: 5px 10px;
	font-size: 1rem;
	cursor: pointer;
}
div.grid {
	display: grid;
	grid-template-rows: auto auto;
	margin-bottom: 1rem;
}
@media (min-width: 800px) {
	div.grid {
		grid-template-rows: auto;
		grid-template-columns: 1fr auto 1fr;
		margin-bottom: 0rem;
		button {
			margin: auto;
		}
		h2 {
			text-align: left;
		}
	}
}
</style>
