<script setup lang="ts">
import Navbar from "./components/Navbar.vue";
</script>

<template>
	<Navbar />
	<RouterView />
</template>

<style>
#app {
	max-width: 90vw;
	margin: auto;
	padding-bottom: 20px;
	padding-top: 20px;
}
</style>
