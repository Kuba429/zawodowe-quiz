<script lang="ts" setup>
const kategorie = ["e12", "e13", "e14", "ee08", "ee09"];
</script>
<template>
	<h1>Strona główna</h1>
	<ul>
		<li><RouterLink to="/pytanie">Wszystkie kategorie</RouterLink></li>
		<li v-for="kat in kategorie">
			<RouterLink :to="'/pytanie/' + kat">
				{{ kat }}
			</RouterLink>
		</li>
	</ul>
</template>
<style></style>
