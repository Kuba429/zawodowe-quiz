<script lang="ts" setup>
import { usePStore } from "../stores/pytanie";
import Odpowiedz from "./Odpowiedz.vue";
const store = usePStore();
</script>
<template>
	<div class="flex">
		{{ store.pytanie?.pytanie }}
		<img
			v-if="store.pytanie?.obrazek"
			:src="store.pytanie.obrazek"
			alt="Obrazek do pytania"
		/>
		<div class="grid">
			<Odpowiedz v-for="odp in store.odpowiedzi" :odpowiedz="odp" />
		</div>
	</div>
</template>
<style scoped lang="scss">
.flex {
	display: flex;
	flex-direction: column;
	background-color: white;
	border-radius: 10px;
	margin: 1rem auto;
	padding: 1rem;
	gap: 1rem;
}
.grid {
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-template-rows: 1fr 1fr;
	gap: 5px;
}
img {
	object-fit: contain;
}

@media (max-width: 400px) {
	.grid {
		grid-template-rows: repeat(4 1fr);
		grid-template-columns: 1fr;
	}
}
</style>
