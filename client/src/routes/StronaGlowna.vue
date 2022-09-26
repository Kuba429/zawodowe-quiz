<script lang="ts" setup>
const kategorie = [
	{
		naglowek: "Ee08",
		link: "/pytanie/ee08",
		zdjecie: "/kategorie-zdjecia/ee08.jpg",
		opis: "Montaż i eksploatacja systemów komputerowych, urządzeń peryferyjnych i sieci",
	},
	{
		naglowek: "Ee09",
		zdjecie: "/kategorie-zdjecia/ee09.jpg",
		link: "/pytanie/ee09",
		opis: "Programowanie, tworzenie i administrowanie stronami internetowymi i bazami danych",
	},
	{
		naglowek: "E12",
		zdjecie: "/kategorie-zdjecia/e12.jpg",
		link: "/pytanie/e12",
		opis: "Montaż i eksploatacja komputerów osobistych oraz urządzeń peryferyjnych",
	},
	{
		naglowek: "E13",
		zdjecie: "/kategorie-zdjecia/e13.jpg",
		link: "/pytanie/e13",
		opis: "Projektowanie lokalnych sieci komputerowych i administrowanie sieciami",
	},
	{
		naglowek: "E14",
		zdjecie: "/kategorie-zdjecia/e14.jpg",
		link: "/pytanie/e14",
		opis: "Tworzenie aplikacji internetowych i baz danych oraz administrowanie bazami",
	},
	{
		naglowek: "Wszystko",
		zdjecie: "/kategorie-zdjecia/wszystko.jpg",
		link: "/pytanie",
		opis: "Pomieszane pytania z każdej kategorii",
	},
];
</script>
<template>
	<h1>Egzamin zawodowy - Technik Informatyk</h1>
	<div class="container">
		<div v-for="k in kategorie">
			<img :src="k.zdjecie" alt="" />
			<h3>{{ k.naglowek }}</h3>
			<p>{{ k.opis }}</p>
			<RouterLink :to="k.link">Losuj Pytanie</RouterLink>
		</div>
	</div>
</template>
<style lang="scss" scoped>
@use "../style/zmienne.scss" as *;
h1 {
	text-align: center;
	color: $niebieski;
}
.container {
	display: flex;
	flex-wrap: wrap;
	flex-basis: 100%;
	justify-content: center;
	gap: 1rem;
	width: 100%;
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
			background-color: $niebieski;
			border-radius: 4px;
			color: white;
			text-decoration: none;
			padding: 5px 15px;
			border: 1px solid $niebieski;
			transition: ease all 0.2s;
			transition-property: color background-color;
			&:hover {
				background-color: white;
				color: $niebieski;
			}
		}
	}
}
</style>
