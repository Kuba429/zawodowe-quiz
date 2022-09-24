<script lang="ts" setup>
import { reactive } from "vue";
import Odpowiedz from "./Odpowiedz.vue";

const { pytanie } = defineProps<{ pytanie: Pytanie }>();
const odpowiedzi = [pytanie.odpA, pytanie.odpB, pytanie.odpC, pytanie.odpD];
const odpStore = reactive({
	poprawna: pytanie.poprawna,
	wybrana: "",
});
</script>
<template>
	<div class="flex">
		{{ pytanie.pytanie }}
		<div class="grid">
			<Odpowiedz
				v-for="odp in odpowiedzi"
				:odpowiedz="odp"
				:odp-store="odpStore"
			/>
		</div>
		{{ odpStore }}
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
