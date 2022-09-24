<script lang="ts" setup>
import { computed } from "vue";
import { usePStore } from "../stores/pytanie";

const { odpowiedz } = defineProps<{
	odpowiedz: string;
}>();
const store = usePStore();
const handleClick = () => {
	if (store.wybranaOdp) return;
	store.$patch({ wybranaOdp: odpowiedz });
};

const klasa = computed(() => ({
	czerwony: store.wybranaOdp === odpowiedz,
	zielony: store.wybranaOdp && store.poprawnaOdp === odpowiedz,
}));
</script>
<template>
	<div :class="klasa" @click="handleClick">
		{{ odpowiedz }}
	</div>
</template>
<style scoped lang="scss">
div {
	background-color: gray;
	&.czerwony {
		background-color: red;
	}
	&.zielony {
		background-color: green;
	}
}
</style>
