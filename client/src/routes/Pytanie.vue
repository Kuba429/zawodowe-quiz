<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const pytanie = ref<Pytanie>();
const jestOk = ref(true);
const { kategoria } = useRoute().params;

onMounted(async () => {
	let url = new URL("http://localhost:3000/pytanie");
	kategoria &&
		(url.search = new URLSearchParams({
			kwal: kategoria as string,
		}).toString());
	try {
		const res = await fetch(url);
		const text = await (await res.blob()).text();
		pytanie.value = JSON.parse(text);
	} catch (error) {
		jestOk.value = false;
		console.error(error);
	}
	return () => {
		console.log("returnasdasd");
	};
});
</script>
<template>
	<div v-if="jestOk">
		<h1>{{ kategoria || "Wszystkie kategorie" }}</h1>
		<pre>{{ pytanie }}</pre>
	</div>
	<div v-else><h1>Nie można połączyć się z serwerem</h1></div>
</template>
<style lang="scss"></style>
