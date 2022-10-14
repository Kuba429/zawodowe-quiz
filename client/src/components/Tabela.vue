<script lang="ts" setup>
import { reactive } from "vue";
import TrescPytania from "./TrescPytania.vue";
import { usePanel } from "../stores/panel";
import { HOST } from "../main";
export type reactiveType = { id: number; element: HTMLElement | undefined }[];

const store = usePanel();
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
		<tr @click="store.ustaw(p)" v-for="p in store.pytania">
			<td>{{ p.id }}</td>
			<TrescPytania :p="p" :pytania-el="pytaniaEl" />
			<td :class="{ poprawna: p.poprawna === 0 }">
				{{ p.odpA }}
			</td>
			<td :class="{ poprawna: p.poprawna === 1 }">
				{{ p.odpB }}
			</td>
			<td :class="{ poprawna: p.poprawna === 2 }">
				{{ p.odpC }}
			</td>
			<td :class="{ poprawna: p.poprawna === 3 }">
				{{ p.odpD }}
			</td>
			<td>{{ ["A", "B", "C", "D"][p.poprawna] }}</td>
			<td>
				<img v-if="p.obrazek" :src="HOST + p.obrazek" />
			</td>
		</tr>
	</table>
</template>
<style lang="scss">
@use "../style/zmienne.scss" as *;
table {
	width: 100%;
}
tr,
td,
th {
	border: 1px solid rgba($niebieski, 0.1);
	text-align: center;
	transition: ease all 0.2s;
}
th {
	color: white;
	background-color: $niebieski;
	padding: 10px;
}
tr {
	&:nth-child(2n) {
		background-color: rgba($color: $niebieski, $alpha: 0.08);
	}
	border: 1px solid rgba($niebieski, 1);
	&:not(:first-child) {
		cursor: pointer;
		&:hover {
			background-color: rgba($niebieski, 0.3);
		}
	}
}
td {
	padding: 10px;
	--wysokosc: 50px;
	height: var(--wysokosc);
	img {
		height: 100%;
		max-height: 100%;
		max-width: 500px;
		object-fit: contain;
	}
}
.poprawna {
	background-color: rgba($color: #00d72f, $alpha: 0.3);
}

// uwaga, mało czytelny css.
// ustawia border-radius dla komórek na rogach tabeli
tr {
	$radius: 8px;
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
