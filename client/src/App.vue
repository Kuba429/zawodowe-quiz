<script setup lang="ts">
import Navbar from "./components/Navbar.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
	<Navbar />
	<div class="route-container">
		<RouterView />
	</div>
	<Footer />
</template>

<style scoped>
.route-container {
	width: 100%;
	max-width: 90vw;
	overflow-x: auto;
	min-height: 100vh;
	margin: 0 auto;
	padding-bottom: 20px;
}
</style>
