<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
			<div class="menu">
				<button class="primary" v-on:click="$router.push('/server')" v-if="is_admin"><i class="fas fa-plus"></i> {{$t("add_new")}}</button>
				<button class="logout-btn" v-on:click="logout"><i class="fas fa-sign-out-alt"></i></button>
			</div>
		</div>
		<server></server>
	</layout>
</template>

<script>
import {layout, server, end} from "../components";

export default {
	components: {layout, server, end},
	methods: {
		logout() {
			this.$store.commit("logout");
			this.$router.push("/login");
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
