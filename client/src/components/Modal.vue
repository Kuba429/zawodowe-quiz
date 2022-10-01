<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { ref } from "vue";
import { useModal } from "../stores/modal";
import ModalForm from "./ModalForm.vue";

const store = useModal();
const { pytanie } = storeToRefs(store);

const klasa = ref<"" | "wyjscie">("");
const zamknij = () => {
	klasa.value = "wyjscie";
	setTimeout(() => {
		store.zamknij();
		klasa.value = "";
		// 300 ms powiazane z czasem animacji
	}, 300);
};
</script>
<template>
	<div v-if="pytanie" @click="zamknij" id="modal-tlo" :class="klasa"></div>
	<ModalForm :zamknij="zamknij" v-if="pytanie" id="modal" :class="klasa" />
</template>
<style scoped lang="scss">
@use "../style/zmienne.scss" as *;
$z-index: 20;
$animacja-czas: 0.3s; // 300ms powiazane z timeoutem

#modal-tlo {
	position: fixed;
	z-index: $z-index;
	top: 0;
	left: 0;
	height: 100%;
	width: 100%;
	background-color: rgba($color: $niebieski, $alpha: 0.2);
	backdrop-filter: blur(5px);
	animation: modal-tlo-in $animacja-czas ease;
	&.wyjscie {
		animation-name: modal-tlo-out;
	}
}
#modal {
	position: fixed;
	z-index: calc($z-index + 5);
	left: 50%;
	top: 50%;
	transform: translate(-50%, -50%);
	background-color: white;
	color: black;
	animation: modal-in $animacja-czas ease;
	&.wyjscie {
		animation-name: modal-out;
	}
}
</style>
