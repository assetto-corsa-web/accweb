<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
			<button><i class="fas fa-plus"></i> {{$t("add_new")}}</button>
		</div>
		<server></server>
	</layout>
</template>

<script>
import {layout, server, end} from "../components";

export default {
	components: {layout, server, end}
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
