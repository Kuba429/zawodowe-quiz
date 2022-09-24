<script lang="ts" setup>
import { usePStore } from "../stores/pytanie";
import Odpowiedz from "./Odpowiedz.vue";
const store = usePStore();
</script>
<template>
	<div class="flex">
		{{ store.pytanie?.pytanie }}
		<div class="grid">
			<Odpowiedz v-for="odp in store.odpowiedzi" :odpowiedz="odp" />
		</div>
	</div>
</template>
<style scoped>
.flex {
	display: flex;
	flex-direction: column;
}
.grid {
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-template-rows: 1fr 1fr;
	gap: 5px;
}
</style>
