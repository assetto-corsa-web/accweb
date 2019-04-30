<template>
	<layout small="true">
		<div class="title">
			<h1>{{$t("title")}}</h1>
		</div>
		<form v-on:submit.prevent="login">
			<label for="password">{{$t("password_label")}}</label>
			<input type="password" id="password" :placeholder="$t('password_placeholder')" v-model="password" />
			<div class="error" v-show="error">{{$t("password_error")}}</div>
			<input class="primary" type="submit" :value="$t('submit_value')" />
		</form>
	</layout>
</template>

<script>
import axios from "axios";
import {layout, server, end} from "../components";

export default {
	components: {layout},
	data() {
		return {
			password: "",
			error: false
		};
	},
	methods: {
		login() {
			axios.post("/api/login", {password: this.password})
			.then(r => {
				this.$store.commit("login", {
					token: r.data.token,
					admin: r.data.admin,
					mod: r.data.mod,
					read_only: r.data.read_only
			});
				this.$router.push("/");
			})
			.catch(() => {
				this.error = true;
			});
		}
	}
}
</script>

<i18n>
{
	"en": {
		"title": "Login",
		"password_label": "Password",
		"password_placeholder": "Password",
		"password_error": "Password incorrect",
		"submit_value": "Login"
	}
}
</i18n>
