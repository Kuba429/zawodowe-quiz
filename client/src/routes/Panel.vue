<script lang="ts" setup>
import { onMounted, onUnmounted } from "vue";
import Modal from "../components/Modal.vue";
import Tabela from "../components/Tabela.vue";
import { usePanel } from "../stores/panel";
import { kategoria, wszyskieKategorie } from "../typy";

const store = usePanel();
const handleInput = async (e: Event) => {
	const kat = (e.target as HTMLSelectElement).value as kategoria;
	store.$patch({ kategoria: kat });
	store.fetchPytania();
};

onMounted(() => {
	window.addEventListener("baza-reset", store.fetchPytania);
	store.fetchPytania();
});
onUnmounted(() => {
	window.removeEventListener("baza-reset", store.fetchPytania);
});
</script>
<template>
	<Modal />
	<div class="grid">
		<label>
			Kategoria:
			<select @input="handleInput">
				<option v-for="k in wszyskieKategorie" :value="k">
					{{ k }}
				</option>
			</select>
		</label>
		<RouterLink to="/dodaj">
			Dodaj pytanie <i class="fa-solid fa-plus"></i>
		</RouterLink>
	</div>
	<Tabela />
</template>
<style scoped lang="scss">
@use "../style/zmienne.scss" as *;
select {
	margin: 20px 0;
	max-width: 100%;
}
.grid {
	display: grid;
	grid-template-columns: 1fr auto 1fr;
	align-items: center;
}
a {
	height: fit-content;
	text-decoration: none;
	color: $niebieski;
	font-weight: bolder;
	padding: 10px 20px;
	border-radius: 5px;
	transition: ease all 0.3s;
	&:hover {
		background-color: $niebieski;
		color: white;
	}
	i {
		margin-left: 5px;
	}
}
@media (max-width: 650px) {
	.grid {
		grid-template-columns: auto;
		margin-top: 10px;
		a {
			justify-self: center;
			grid-row: 1;
		}
	}
}
</style>
