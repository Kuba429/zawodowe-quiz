<script lang="ts" setup>
import { onMounted, ref } from "vue";
import Modal from "../components/Modal.vue";
import Tabela from "../components/Tabela.vue";
import { useModal } from "../stores/modal";
import type { kategoria, Pytanie } from "../typy";

const wszyskieKategorie: kategoria[] = ["e12", "e13", "e14", "ee08", "ee09"];

const pytania = ref<Pytanie[]>([]);
const store = useModal();
const handleInput = async (e: Event) => {
	const kat = (e.target as HTMLSelectElement).value as kategoria;
	store.$patch({ kategoria: kat });
	fetchPytania(kat);
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
	<Modal />
	<label>
		Kategoria:
		<select @input="handleInput">
			<option v-for="k in wszyskieKategorie" :value="k">{{ k }}</option>
		</select>
	</label>
	<Tabela :pytania="pytania" />
</template>
<style scoped lang="scss">
select {
	margin: 20px 0;
	max-width: 100%;
}
</style>
