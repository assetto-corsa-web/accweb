<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
			<div>
				<button class="new-server"><i class="fas fa-plus"></i> {{$t("add_new")}}</button>
				<button v-on:click="logout"><i class="fas fa-sign-out-alt"></i></button>
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

<i18n>
{
	"en": {
		"title": "Servers",
		"add_new": "Add Server"
	}
}
</i18n>
