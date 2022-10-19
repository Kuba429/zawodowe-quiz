<script setup lang="ts">
import { HOST } from "../main";

const handleReset = async () => {
	const res = await fetch(`${HOST}/baza/reset`);
	if (!res.ok) {
		alert("Nie można zresetować bazy");
		return;
	}
	alert("Baza zresetowana");
	window.dispatchEvent(new Event("baza-reset"));
};
</script>
<template>
	<footer>
		<a href="/dok.zip" download
			><i class="fa-regular fa-file"></i>Dokumentacja</a
		>
		<a href="https://github.com/Kuba429/zawodowe-quiz" target="_blank">
			<i class="fa-brands fa-github"></i> Github
		</a>
		<span @click="handleReset">
			<i class="fa-regular fa-database"></i>Resetuj bazę danych
		</span>
	</footer>
</template>
<style lang="scss">
footer {
	background: #092442;
	color: white;
	margin-top: auto;
	padding: 30px 10px;
	text-align: center;
	display: flex;
	justify-content: space-evenly;
	a,
	span {
		color: white;
		text-decoration: none;
		cursor: pointer;
		transition: ease 0.2s color;
		&:hover {
			color: #9cc6f3;
		}
	}
	i {
		padding-right: 8px;
	}
}
</style>
