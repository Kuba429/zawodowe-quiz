<script lang="ts" setup>
import { Pytanie } from "../typy";

const { pytania } = defineProps<{ pytania: Pytanie[] }>();
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
			<td>{{ p.id }}</td>
			<td>
				<div>{{ p.pytanie }}</div>
			</td>
			<td class="p">{{ p.odpA }}</td>
			<td class="p">{{ p.odpB }}</td>
			<td class="p">{{ p.odpC }}</td>
			<td class="p">{{ p.odpD }}</td>
			<td>{{ p.poprawna }}</td>
			<td>
				<a target="_blank" :href="p.obrazek">{{ p.obrazek }}</a>
			</td>
		</tr>
	</table>
</template>
<style lang="scss">
@use "../style/zmienne.scss" as *;
table,
tr,
td,
th {
	border: 1px solid rgba($niebieski, 0.4);
	text-align: center;
	& > div {
		text-align: left;
		width: 100%;
		height: 100px;
		resize: both;
		overflow: auto;
	}
}
.p {
	padding: 10px;
}
</style>
