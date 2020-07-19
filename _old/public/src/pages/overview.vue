<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
			<div class="menu">
				<button class="primary" v-on:click="$router.push('/server')" v-if="is_admin"><i class="fas fa-plus"></i> {{$t("add_new")}}</button>
				<button v-on:click="loadServer(true)"><i class="fas fa-sync"></i> {{$t("refresh")}}</button>
				<button class="logout-btn" v-on:click="logout"><i class="fas fa-sign-out-alt"></i></button>
			</div>
		</div>
		<server v-for="s in server"
			:key="s.id"
			:server="s"
			:ro="!is_mod && !is_admin"
			v-on:copied="loadServer"
			v-on:deleted="loadServer"
			v-on:started="loadServer"
			v-on:stopped="loadServer"></server>
		<p v-if="!server || !server.length">No servers found.</p>
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
		this.refreshList();
	},
	methods: {
		logout() {
			this.$store.commit("logout");
			this.$router.push("/login");
		},
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
				})
				.catch(e => {
					this.$store.commit("toast", this.$t("receive_server_list_error"))
				});
			}, timeout);
		},
		refreshList() {
			this.loadServer();
			setTimeout(() => {
				this.refreshList();
			}, 5000);
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
		"add_new": "Add Server",
		"refresh": "Refresh",
		"receive_server_list_error": "Error receiving server list."
	}
}
</i18n>
