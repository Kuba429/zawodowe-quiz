<script lang="ts" setup>
const kategorie = [
	{
		nazwa: "ee08",
		opis: "Montaż i eksploatacja systemów komputerowych, urządzeń peryferyjnych i sieci",
	},
	{
		nazwa: "ee09",
		opis: "Programowanie, tworzenie i administrowanie stronami internetowymi i bazami danych",
	},
	{
		nazwa: "e12",
		opis: "Montaż i eksploatacja komputerów osobistych oraz urządzeń peryferyjnych",
	},
	{
		nazwa: "e13",
		opis: "Projektowanie lokalnych sieci komputerowych i administrowanie sieciami",
	},
	{
		nazwa: "e14",
		opis: "Tworzenie aplikacji internetowych i baz danych oraz administrowanie bazami",
	},
];
</script>
<template>
	<h1>Egzamin zawodowy - Technik Informatyk</h1>
	<div class="container">
		<div v-for="k in kategorie">
			<img :src="`/kategorie-zdjecia/${k.nazwa}.jpg`" alt="" />
			<h3>{{ k.nazwa }}</h3>
			<p>{{ k.opis }}</p>
			<RouterLink :to="'/pytanie/' + k.nazwa">Losuj Pytanie</RouterLink>
		</div>
	</div>
</template>
<style lang="scss" scoped>
@use "../style/zmienne.scss" as *;
h1 {
	text-align: center;
}
.container {
	display: flex;
	flex-wrap: wrap;
	flex-basis: 100%;
	justify-content: center;
	gap: 1rem;
	max-width: 90vw;
	margin: auto;
	padding-bottom: 50px;
	& > div {
		width: 300px;
		text-align: center;
		background: white;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		border-radius: 8px;
		overflow: hidden;
		img {
			width: 100%;
			object-fit: cover;
		}
		h3 {
			text-transform: capitalize;
			margin: 10px;
		}
		p {
			margin: 0;
			padding: 0 10px;
			flex-grow: 1;
		}
		a {
			margin: 10px;
			background-color: $niebeski;
			border-radius: 4px;
			color: white;
			text-decoration: none;
			padding: 5px 15px;
		}
	}
}
</style>
