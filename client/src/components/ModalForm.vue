<script lang="ts" setup>
import { reactive } from "vue";
import { useModal } from "../stores/modal";

const { zamknij } = defineProps<{ zamknij: () => void }>();
const store = useModal();
const pytanie = reactive({ ...store.pytanie });

const handleSubmit = async (e: Event) => {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	formData.set("kategoria", store.kategoria);
	if ((formData.get("obrazek") as File)?.size === 0) {
		formData.set("obrazek", "");
	} else if (!formData.has("obrazek")) {
		formData.set("obrazek", pytanie.obrazek || "");
	}
	const res = await fetch("http://localhost:3000/update-pytanie", {
		method: "post",
		body: formData,
	});
	if (res.ok) zamknij();
	// todo handle error
};
</script>
<template>
	<form @submit="handleSubmit">
		<label>
			Tresc:
			<textarea name="pytanie" :value="pytanie.pytanie" type="text" />
		</label>
		<label>
			Odp A: <input name="odpA" :value="pytanie.odpA" type="text" />
		</label>
		<label>
			Odp B: <input name="odpB" :value="pytanie.odpB" type="text" />
		</label>
		<label>
			Odp C: <input name="odpC" :value="pytanie.odpC" type="text" />
		</label>
		<label>
			Odp D: <input name="odpD" :value="pytanie.odpD" type="text" />
		</label>
		<label>
			Poprawna odpowiedź:
			<select name="poprawna">
				<option value="0">A</option>
				<option value="1">B</option>
				<option value="2">C</option>
				<option value="3">D</option>
			</select>
		</label>
		<label v-if="pytanie.obrazek">
			<button @click="pytanie.obrazek = ''">usun</button>
		</label>
		<label v-else> Obrazek: <input name="obrazek" type="file" /> </label>
		<button type="submit">submit</button>
	</form>
</template>
<style scoped lang="scss">
form {
	width: 80vw;
	display: flex;
	gap: 1rem;
	flex-direction: column;
	overflow: hidden;
	padding: 2rem;
	textarea {
		width: 100%;
		max-width: 100%;
	}
}
</style>
