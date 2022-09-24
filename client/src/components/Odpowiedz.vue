<script lang="ts" setup>
import { computed, reactive } from "vue";

const { odpowiedz, odpStore } = defineProps<{
	odpowiedz: string;
	odpStore: {
		poprawna: string;
		wybrana: string;
	};
}>();

const klasa = computed(() => ({
	czerwony: odpStore.wybrana === odpowiedz,
	zielony: odpStore.wybrana && odpowiedz === odpStore.poprawna,
}));

const handleClick = () => {
	if (odpStore.wybrana) return;
	odpStore.wybrana = odpowiedz;
};
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
