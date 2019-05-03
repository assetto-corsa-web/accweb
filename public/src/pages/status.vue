<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
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
		loadServer() {
			axios.get("/api/server")
			.then(r => {
				this.server = r.data;
			});
		}
	}
}
</script>

<i18n>
{
	"en": {
		"title": "Server Status"
	}
}
</i18n>
