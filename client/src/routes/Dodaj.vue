<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { statusType, wszyskieKategorie } from "../typy";

const router = useRouter();
const status = ref<statusType>("sukces");
const handleSubmit = async (e: Event) => {
	status.value = "ladowanie";
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	try {
		const res = await fetch("http://localhost:3000/dodaj-pytanie", {
			method: "post",
			body: formData,
		});
		if (!res.ok) throw new Error(await res.text());
		status.value = "sukces";
		router.push("/panel");
	} catch (err) {
		status.value = "blad";
		console.error(err);
	}
};
</script>

<template>
	<h1>Dodaj pytanie</h1>
	<form @submit="handleSubmit">
		<label>
			Tresc:
			<textarea required rows="3" name="pytanie" type="text" />
		</label>
		<label
			>Kategoria:
			<select name="kategoria">
				<option v-for="kat in wszyskieKategorie" :value="kat">
					{{ kat }}
				</option>
			</select></label
		>
		<label>Odp A: <input required name="odpA" type="text" /></label>
		<label>Odp B: <input required name="odpB" type="text" /></label>
		<label>Odp C: <input required name="odpC" type="text" /></label>
		<label>Odp D: <input required name="odpD" type="text" /></label>
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
		<button v-if="status === 'sukces'" type="submit">Zatwierdź</button>
		<button v-else-if="status === 'blad'" type="submit">
			Spróbuj ponownie
		</button>
		<button v-else disabled type="submit">Wysyłam</button>
	</form>
</template>

<style lang="scss" scoped>
form {
	display: flex;
	flex-direction: column;
}
</style>
