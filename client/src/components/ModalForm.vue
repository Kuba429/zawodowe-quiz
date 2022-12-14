<script lang="ts" setup>
import { reactive, ref } from "vue";
import { API, HOST } from "../main";
import { usePanel } from "../stores/panel";
import { Pytanie, statusType } from "../typy";

const { zamknij } = defineProps<{ zamknij: () => void }>();
const store = usePanel();
const pytanie = reactive({ ...store.modalPytanie });
const status = ref<statusType>("sukces");
const handleSubmit = async (e: Event) => {
	e.preventDefault();
	status.value = "ladowanie";
	const formData = new FormData(e.target as HTMLFormElement);
	formData.set("kategoria", store.kategoria);
	formData.set("id", pytanie.id!.toString());
	if ((formData.get("obrazek") as File)?.size === 0) {
		formData.set("obrazek", "");
	} else if (!formData.has("obrazek")) {
		formData.set("obrazek", pytanie.obrazek || "");
	}
	try {
		const res = await fetch(`${API}/update-pytanie`, {
			method: "post",
			body: formData,
		});
		if (!res.ok) {
			status.value = "blad";
			return;
		}
		status.value = "sukces";
		store.update((await res.json()) as Pytanie);
		zamknij();
	} catch (error) {
		status.value = "blad";
	}
};

const podgladZdjecia = ref<string>();
const handleZdjecieInput = async (e: Event) => {
	const x = (e.target as HTMLInputElement).files?.item(0);
	if (!x) return;
	podgladZdjecia.value = window.URL.createObjectURL(x);
};

const handleUsun = async () => {
	if (!pytanie.id || !confirm("Czy na pewno chcesz usunąć pytanie?")) return;
	const url = new URL(`${API}/usun-pytanie`);
	url.search = new URLSearchParams({
		kategoria: store.kategoria,
		idPytania: pytanie.id.toString(),
	}).toString();
	try {
		const res = await fetch(url);
		if (res.ok) {
			store.usun(pytanie.id);
			zamknij();
		}
	} catch (error) {
		console.error(error);
	}
};
</script>
<template>
	<form @submit="handleSubmit">
		<label>
			Tresc:
			<textarea
				rows="3"
				name="pytanie"
				:value="pytanie.pytanie"
				type="text"
			/>
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
					v-for="litera in ['a', 'b', 'c', 'd']"
					:selected="pytanie.poprawna === litera"
					:value="litera"
				>
					{{ litera.toUpperCase() }}
				</option>
			</select>
		</label>
		<div class="obrazek-wrapper" v-if="pytanie.obrazek">
			Obrazek:
			<img :src="HOST + pytanie.obrazek" alt="" />
			<button @click="pytanie.obrazek = ''">
				<i class="brak-margin fa-solid fa-delete-left"></i>
			</button>
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
		<div class="button-wrapper">
			<button class="usun-pytanie" @click="handleUsun" type="button">
				Usun <i class="fa-regular fa-trash-can"></i>
			</button>
			<button v-if="status === 'sukces'" type="submit">
				Zatwierdź <i class="fa-solid fa-check"></i>
			</button>
			<button v-else-if="status === 'blad'" type="submit">
				spróbuj ponownie
			</button>
			<button v-else-if="status === 'ladowanie'" disabled type="submit">
				Wysyłam
			</button>
		</div>
	</form>
</template>
<style scoped lang="scss">
@use "../style/zmienne.scss" as *;
form {
	width: 60vw;
	max-height: 90%;
	display: flex;
	gap: 1rem;
	flex-direction: column;
	overflow: hidden;
	overflow-y: auto;
	box-sizing: border-box;
	padding: 2rem;
	border-radius: 10px;
	font-size: 1rem;
	.button-wrapper {
		display: flex;
		gap: 2px;
		width: 100%;
		& > button {
			flex-grow: 1;
			margin: 0;
		}
		& > .usun-pytanie {
			background-color: rgba($color: red, $alpha: 0.7);
			color: white;
			border: 1px solid transparent;
			&:hover {
				background-color: rgba($color: red, $alpha: 0.6);
				border-color: red;
			}
		}
		& > button:nth-child(2) {
			flex-grow: 5;
			margin-left: auto;
		}
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
			padding: 5px 10px 5px 8px;
			&:hover {
				background-color: rgba($color: red, $alpha: 0.1);
				border-radius: 2px;
			}
		}
	}
}
i:not(.brak-margin) {
	margin-left: 5px;
}
@media (max-width: 520px) {
	form {
		width: 98vw;
		max-width: 98vw;
	}
}
</style>
