<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
			<div class="menu">
				<button class="primary" v-on:click="$router.push('/server')" v-if="is_admin"><i class="fas fa-plus"></i> {{$t("add_new")}}</button>
				<button class="logout-btn" v-on:click="logout"><i class="fas fa-sign-out-alt"></i></button>
			</div>
		</div>
		<server v-for="s in server" :server="s"></server>
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
		logout() {
			this.$store.commit("logout");
			this.$router.push("/login");
		},
		loadServer() {
			axios.get("/api/server")
			.then(r => {
				this.server = r.data;
			});
		}
	}
}
</script>

<style scoped>
.logout-btn .fas {
	margin: 0;
}
</style>

<i18n>
{
	"en": {
		"title": "Servers",
		"add_new": "Add Server"
	}
}
</i18n>
