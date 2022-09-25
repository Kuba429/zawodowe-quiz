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
	<h2>{{ kategoria || "Wszystkie kategorie" }}</h2>
	<Pytanie v-if="store.pytanie && store.status === 'success'" />
	<button class="nowe-pytanie" @click="store.nowePytanie">
		Nowe pytanie
	</button>
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
</style>
