<script setup lang="ts">
import { wszyskieKategorie } from "../typy";
const handleSubmit = async (e: Event) => {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const res = await fetch("http://localhost:3000/dodaj-pytanie", {
		method: "post",
		body: formData,
	});
	console.log(res);
};
</script>

<template>
	<h1>Dodaj pytanie</h1>
	<form @submit="handleSubmit">
		<label>
			Tresc:
			<textarea rows="3" name="pytanie" type="text" />
		</label>
		<label
			>Kategoria:
			<select name="kategoria">
				<option v-for="kat in wszyskieKategorie" :value="kat">
					{{ kat }}
				</option>
			</select></label
		>
		<label>Odp A: <input name="odpA" type="text" /></label>
		<label>Odp B: <input name="odpB" type="text" /></label>
		<label>Odp C: <input name="odpC" type="text" /></label>
		<label>Odp D: <input name="odpD" type="text" /></label>
		<label>
			Poprawna odpowiedź:
			<select name="poprawna">
				<option v-for="n in [0, 1, 2, 3]" :value="n">
					{{ ["A", "B", "C", "D"][n] }}
				</option>
			</select>
		</label>
		<label>
			Obrazek: <input type="file" name="obrazek" accept="image/*" />
		</label>
		<button type="submit">Zatwierdź</button>
	</form>
</template>

<style lang="scss" scoped>
form {
	display: flex;
	flex-direction: column;
}
label {
	display: flex;
	flex-wrap: wrap;
}
input,
textarea {
	flex-grow: 1;
}
textarea {
	width: 100%;
	max-width: 100%;
}
</style>
