<script lang="ts" setup>
import { reactive, ref } from "vue";
import { usePanel } from "../stores/panel";
import { Pytanie } from "../typy";

const { zamknij } = defineProps<{ zamknij: () => void }>();
const store = usePanel();
const pytanie = reactive({ ...store.modalPytanie });

const handleSubmit = async (e: Event) => {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	formData.set("kategoria", store.kategoria);
	formData.set("id", pytanie.id!.toString());
	if ((formData.get("obrazek") as File)?.size === 0) {
		formData.set("obrazek", "");
	} else if (!formData.has("obrazek")) {
		formData.set("obrazek", pytanie.obrazek || "");
	}
	const res = await fetch("http://localhost:3000/update-pytanie", {
		method: "post",
		body: formData,
	});
	if (res.ok) {
		store.update((await res.json()) as Pytanie);
		zamknij();
	}
	// todo handle error
};

const podgladZdjecia = ref<string>();
const handleZdjecieInput = async (e: Event) => {
	const x = (e.target as HTMLInputElement).files?.item(0);
	if (!x) return;
	podgladZdjecia.value = window.URL.createObjectURL(x);
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
				<option
					v-for="n in [0, 1, 2, 3]"
					:selected="pytanie.poprawna === n"
					:value="n"
				>
					{{ ["A", "B", "C", "D"][n] }}
				</option>
			</select>
		</label>
		<div class="obrazek-wrapper" v-if="pytanie.obrazek">
			Obrazek:
			<img :src="'http://localhost:3000' + pytanie.obrazek" alt="" />
			<button @click="pytanie.obrazek = ''">usun</button>
		</div>
		<label class="obrazek-wrapper" v-else>
			Obrazek:
			<img
				v-if="podgladZdjecia"
				:src="podgladZdjecia"
				alt="podgląd zdjęcia"
			/>
			<input
				@input="handleZdjecieInput"
				name="obrazek"
				type="file"
				accept="image/*"
			/>
		</label>
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
	.obrazek-wrapper {
		display: flex;
		align-items: center;
		gap: 5px;

		img {
			display: inline;
			object-fit: contain;
			height: 1em;
			max-width: 70%;
		}
		button {
			border: none;
			background-color: transparent;
			color: red;
			cursor: pointer;
			&:hover {
				background-color: rgba($color: red, $alpha: 0.1);
				border-radius: 2px;
			}
		}
	}
}
</style>
