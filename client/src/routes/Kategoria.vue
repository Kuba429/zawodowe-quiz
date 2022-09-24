<script lang="ts" setup>
import { onMounted } from "vue";
import { useRoute } from "vue-router";
import Pytanie from "../components/Pytanie.vue";
import { usePStore } from "../stores/pytanie";

const store = usePStore();
const { kategoria } = useRoute().params;
onMounted(() => {
	store.$patch({ kategoria: kategoria as string });
	store.nowePytanie();
});
</script>
<template>
	<h1>{{ kategoria || "Wszystkie kategorie" }}</h1>
	<button @click="store.nowePytanie">Losuj nowe pytanie</button>
	<Pytanie v-if="store.pytanie && store.status === 'success'" />
</template>
<style lang="scss"></style>
