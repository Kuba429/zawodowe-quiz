<script lang="ts" setup>
import { onMounted, ref } from "vue";
import Modal from "../components/Modal.vue";
import Tabela from "../components/Tabela.vue";
import { usePanel } from "../stores/panel";
import type { kategoria, Pytanie } from "../typy";

const wszyskieKategorie: kategoria[] = ["e12", "e13", "e14", "ee08", "ee09"];

const store = usePanel();
const handleInput = async (e: Event) => {
	const kat = (e.target as HTMLSelectElement).value as kategoria;
	store.$patch({ kategoria: kat });
	store.fetchPytania();
};

onMounted(() => store.fetchPytania());
</script>
<template>
	<Modal />
	<label>
		Kategoria:
		<select @input="handleInput">
			<option v-for="k in wszyskieKategorie" :value="k">{{ k }}</option>
		</select>
	</label>
	<Tabela />
</template>
<style scoped lang="scss">
select {
	margin: 20px 0;
	max-width: 100%;
}
</style>
