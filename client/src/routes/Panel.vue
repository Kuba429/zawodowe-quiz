<script lang="ts" setup>
import { onMounted, ref } from "vue";
import type { kategoria, Pytanie } from "../typy";
import Tabela from "../components/Tabela.vue";

const wszyskieKategorie: kategoria[] = ["e12", "e13", "e14", "ee08", "ee09"];

const pytania = ref<Pytanie[]>([]);

const handleInput = async (e: Event) => {
	fetchPytania((e.target as HTMLSelectElement).value as kategoria);
};
const fetchPytania = async (kat: kategoria) => {
	const url = new URL("http://localhost:3000/wszystkie-pytania");
	url.search = new URLSearchParams({
		kwal: kat,
	}).toString();
	try {
		const res = await fetch(url);
		const text = await (await res.blob()).text();
		const pytaniaRaw: Pytanie[] = JSON.parse(text);
		pytaniaRaw.forEach(
			(p) =>
				p.obrazek && (p.obrazek = "http://localhost:3000" + p.obrazek)
		);
		pytania.value = pytaniaRaw;
	} catch (error) {
		console.error(error);
	}
};
onMounted(() => fetchPytania(wszyskieKategorie[0]));
</script>
<template>
	<select @input="handleInput">
		<option v-for="k in wszyskieKategorie" :value="k">{{ k }}</option>
	</select>
	<Tabela :pytania="pytania" />
</template>
<style lang="scss"></style>
