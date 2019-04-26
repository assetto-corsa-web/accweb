<template>
	<layout small="true">
		<div class="title">
			<h1>{{$t("title")}}</h1>
		</div>
		<form v-on:submit.prevent="login">
			<label for="password">{{$t("password_label")}}</label>
			<input type="password" id="password" :placeholder="$t('password_placeholder')" v-model="password" />
			<div class="error">{{$t("password_error")}}</div>
			<input type="submit" :value="$t('submit_value')" />
		</form>
	</layout>
</template>

<script>
import {layout, server, end} from "../components";

export default {
	components: {layout},
	data() {
		return {
			password: ""
		};
	},
	methods: {
		login() {
			console.log(this.password);
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
