<template>
	<layout small="true">
		<div class="title">
			<h1>{{$t("title")}}</h1>
		</div>
		<form v-on:submit.prevent="login">
			<field type="password" :label="$t('password_label')" :placeholder="$t('password_placeholder')" :error="error" v-model="password"></field>
			<input class="primary" type="submit" :value="$t('submit_value')" />
		</form>
	</layout>
</template>

<script>
import axios from "axios";
import {layout, field} from "../components";

export default {
	components: {layout, field},
	data() {
		return {
			password: "",
			error: ""
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
				this.error = this.$t("password_error");
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
