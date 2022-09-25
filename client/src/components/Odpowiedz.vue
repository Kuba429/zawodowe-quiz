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
	active: !store.wybranaOdp,
	czerwony: store.wybranaOdp === odpowiedz,
	zielony: store.wybranaOdp && store.poprawnaOdp === odpowiedz,
}));
</script>
<template>
	<button :class="klasa" @click="handleClick">
		{{ odpowiedz }}
	</button>
</template>
<style scoped lang="scss">
@use "../style/zmienne.scss" as *;
button {
	padding: 10px;
	border-radius: 7px;
	border: solid 1px $niebieski;
	color: $niebieski;
	background-color: white;
	font-size: 1rem;
	text-align: left;
	cursor: pointer;
	&.czerwony {
		$kolor: rgb(255, 72, 72);
		background-color: $kolor;
		border-color: $kolor;
		color: white;
	}
	&.zielony {
		$kolor: rgb(0, 208, 0);
		background-color: $kolor;
		border-color: $kolor;
		color: white;
	}
	transition: ease all 0.2s;
	&.active:active {
		transform: scale(98%);
	}
}
</style>
