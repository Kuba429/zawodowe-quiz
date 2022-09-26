<script lang="ts" setup>
import { onMounted, ref } from "vue";
const nav = ref<HTMLElement>();
onMounted(() => {
	if (!nav.value) return;
	const wysokoscNav = nav.value.getBoundingClientRect().height;

	document.querySelector<HTMLElement>("#app")!.style.marginTop =
		wysokoscNav + "px";
});
</script>
<template>
	<div class="nav-wrapper">
		<nav ref="nav">
			<RouterLink to="/">Strona Główna</RouterLink>
			<div class="spacer"></div>
			<RouterLink to="/pytanie">Pytanie</RouterLink>
			<RouterLink to="/">Panel</RouterLink>
		</nav>
	</div>
</template>
<style scoped lang="scss">
@use "../style/zmienne.scss" as *;

.nav-wrapper {
	background-color: white;
	color: $niebieski;
	box-shadow: 0px 0px 3px 0px $niebieski;
	font-size: 1.2rem;
	padding: 10px;
	box-sizing: content-box;
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
	nav {
		width: 90%;
		margin: auto;
		display: flex;
		gap: 1rem;
		justify-content: space-between;
		.spacer {
			flex-grow: 1;
		}
		a {
			color: inherit;
			text-decoration: none;
		}
	}
}
</style>
