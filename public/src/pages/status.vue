<template>
	<layout>
		<div class="title">
			<h1>{{$t("title")}}</h1>
		</div>
		<server ro="true"></server>
	</layout>
</template>

<script>
import {layout, server, end} from "../components";

export default {
	components: {layout, server, end},
	methods: {
		
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
