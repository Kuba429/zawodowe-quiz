<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import Pytanie from "../components/Pytanie.vue";

const pytanie = ref<Pytanie>();
const error = ref<string>();
const { kategoria } = useRoute().params;

const pobierzNowePytanie = async () => {
	let url = new URL("http://localhost:3000/pytanie");
	kategoria &&
		(url.search = new URLSearchParams({
			kwal: kategoria as string,
		}).toString());
	try {
		const res = await fetch(url);
		const text = await (await res.blob()).text();
		pytanie.value = JSON.parse(text);
	} catch (err) {
		error.value = err as string;
		console.error(err);
	}
};
onMounted(pobierzNowePytanie);
</script>
<template>
	<h1>{{ kategoria || "Wszystkie kategorie" }}</h1>
	<Pytanie v-if="pytanie" :pytanie="pytanie" />
	<div v-else-if="!error">Ładowanie...</div>
	<div v-else><h1>Nie można połączyć się z serwerem</h1></div>
</template>
<style lang="scss"></style>
