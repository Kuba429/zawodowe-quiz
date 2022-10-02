<script lang="ts" setup>
import { onMounted, ref } from "vue";
import type { Pytanie } from "../typy";
import type { reactiveType } from "./Tabela.vue";

const { pytaniaEl, p } = defineProps<{
	pytaniaEl: reactiveType;
	p: Pytanie;
}>();
const elRef = ref<HTMLElement>();
const observer = new ResizeObserver(() => {
	pytaniaEl.forEach((x) => {
		if (!x.element || !elRef.value) return;
		x.element.style.width = elRef.value.style.width;
	});
});
onMounted(() => {
	pytaniaEl.push({ element: elRef.value, id: p.id });
	observer.observe(elRef.value!);
});
</script>
<template>
	<td>
		<div ref="elRef">{{ p.pytanie }}</div>
	</td>
</template>
<style scoped lang="scss">
td {
	padding: 0;
}
div {
	text-align: left;
	width: 120px;
	height: 100px;
	min-width: 100px;
	min-height: var(--wysokosc);
	resize: both;
	overflow: auto;
	padding: 10px;
}
</style>
