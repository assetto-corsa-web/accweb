<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
			<div class="menu">
				<button v-on:click="loadServer(true)"><i class="fas fa-sync"></i> {{$t("refresh")}}</button>
			</div>
		</div>
		<server v-for="s in server" :server="s" ro="true"></server>
	</layout>
</template>

<script>
import axios from "axios";
import {layout, server, end} from "../components";

export default {
	components: {layout, server, end},
	data() {
		return {
			server: []
		};
	},
	mounted() {
		this.loadServer();
	},
	methods: {
		loadServer(refresh) {
			let timeout = 0;

			if(refresh) {
				this.server = [];
				timeout = 100;
			}

			setTimeout(() => {
				axios.get("/api/server")
				.then(r => {
					this.server = r.data;
				});
			}, timeout);
		}
	}
}
</script>

<i18n>
{
	"en": {
		"title": "Server Status",
		"refresh": "Refresh"
	}
}
</i18n>
