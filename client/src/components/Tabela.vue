<script lang="ts" setup>
import { reactive } from "vue";
import { Pytanie } from "../typy";
import TrescPytania from "./TrescPytania.vue";
export type reactiveType = { id: number; element: HTMLElement | undefined }[];

const { pytania } = defineProps<{ pytania: Pytanie[] }>();
const pytaniaEl = reactive<reactiveType>([]);
</script>
<template>
	<table>
		<tr>
			<th>ID</th>
			<th>Treść</th>
			<th>Odpowiedź A</th>
			<th>Odpowiedź B</th>
			<th>Odpowiedź C</th>
			<th>Odpowiedź D</th>
			<th>Poprawna odpowiedź</th>
			<th>Obraz</th>
		</tr>
		<tr v-for="p in pytania">
			<td class="p">{{ p.id }}</td>
			<TrescPytania :p="p" :pytania-el="pytaniaEl" />
			<td :class="{ p: true, poprawna: p.poprawna === 0 }">
				{{ p.odpA }}
			</td>
			<td :class="{ p: true, poprawna: p.poprawna === 1 }">
				{{ p.odpB }}
			</td>
			<td :class="{ p: true, poprawna: p.poprawna === 2 }">
				{{ p.odpC }}
			</td>
			<td :class="{ p: true, poprawna: p.poprawna === 3 }">
				{{ p.odpD }}
			</td>
			<td>{{ ["A", "B", "C", "D"][p.poprawna] }}</td>
			<td>
				<a target="_blank" :href="p.obrazek">{{ p.obrazek }}</a>
			</td>
		</tr>
	</table>
</template>
<style lang="scss">
@use "../style/zmienne.scss" as *;
table {
	width: 100%;
	padding-right: 5vw;
}
tr,
td,
th {
	border: 1px solid rgba($niebieski, 0.4);
	text-align: center;
}
th {
	color: white;
	background-color: $niebieski;
}
.p,
th {
	padding: 10px;
}

.poprawna {
	background-color: rgba($color: #00d72f, $alpha: 0.3);
}

// uwaga, mało czytelny css.
// ustawia border-radius dla komórek na rogach tabeli
tr {
	$radius: 5px;
	&:first-child {
		th:first-child {
			border-top-left-radius: $radius;
		}
		th:last-child {
			border-top-right-radius: $radius;
		}
	}
	&:last-child {
		td:first-child {
			border-bottom-left-radius: $radius;
		}
		td:last-child {
			border-bottom-right-radius: $radius;
		}
	}
}
</style>
