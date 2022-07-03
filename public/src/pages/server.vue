<template>
	<layout>
		<div class="title">
			<h1>{{ servername }}</h1>
			<div class="menu">
				<button class="primary" v-on:click="save" v-if="is_admin">
					<i class="fas fa-save"></i> {{ $t("save") }}
				</button>
				<button v-on:click="$router.push('/')" v-if="is_admin">
					<i class="fas fa-ban"></i> {{ $t("cancel") }}
				</button>
				<button class="primary" v-on:click="$router.push('/')" v-if="!is_admin">
					<i class="fas fa-arrow-left"></i> {{ $t("back") }}
				</button>
			</div>
		</div>
		<div>
			<tabs :options="{ useUrlFragment: false }">
				<tab name="Basic Settings">
					<field :label="$t('servername_label')" v-model="servername"></field>
					<accweb ref="accweb"></accweb>
					<basic ref="basic"></basic>
				</tab>
				<tab name="Advanced Settings">
					<settings ref="settings"></settings>
				</tab>
				<tab name="Event">
					<event ref="event"></event>
				</tab>
				<tab name="eventrules">
					<eventrules ref="eventrules"></eventrules>
				</tab>
				<tab name="entrylist">
					<entrylist ref="entrylist"></entrylist>
				</tab>
				<tab name="bop">
					<bop ref="bop"></bop>
				</tab>
				<tab name="assistrules">
					<assistrules ref="assistrules"></assistrules>
				</tab>
			</tabs>
		</div>
	</layout>
</template>

<script>
import axios from "axios";
import {
	layout,
	end,
	accweb,
	basic,
	settings,
	event,
	eventrules,
	entrylist,
	bop,
	assistrules,
	field
} from "../components";

export default {
	components: {
		layout,
		end,
		accweb,
		basic,
		settings,
		event,
		eventrules,
		entrylist,
		bop,
		assistrules,
		field
	},
	data() {
		return {
			id: 0,
			servername: "Server name (by accweb)",
			configurationJson: null,
			settingsJson: null,
			eventJson: null,
			eventRulesJson: null,
			entrylistJson: null,
			bopJson: null,
			assistRulesJson: null
		};
	},
	mounted() {
		this.id = this.$route.query.id;

		if (this.id) {
			this.loadServer();
		}
	},
	methods: {
		loadServer() {
			axios.get("/api/instance/" + this.id).then(r => {
				let settings = r.data.acc.settings;
				settings.passwordIsEmpty = r.data.accExtraSettings.passwordIsEmpty;
				settings.adminPasswordIsEmpty =
					r.data.accExtraSettings.adminPasswordIsEmpty;
				settings.spectatorPasswordIsEmpty =
					r.data.accExtraSettings.spectatorPasswordIsEmpty;

				this.servername = r.data.acc.settings.serverName;
				this.$refs.accweb.setData(r.data.accWeb);
				this.$refs.basic.setData(r.data.acc.configuration);
				this.$refs.settings.setData(r.data.acc.settings);
				this.$refs.event.setData(r.data.acc.event);
				this.$refs.eventrules.setData(r.data.acc.eventRules);
				this.$refs.entrylist.setData(r.data.acc.entrylist);
				this.$refs.bop.setData(r.data.acc.bop);
				this.$refs.assistrules.setData(r.data.acc.assistRules);
			});
		},
		save() {
			let accWeb = this.$refs.accweb.getData();
			let configuration = this.$refs.basic.getData();
			let settings = this.$refs.settings.getData();
			let event = this.$refs.event.getData();
			let eventRules = this.$refs.eventrules.getData();
			let entrylist = this.$refs.entrylist.getData();
			let bop = this.$refs.bop.getData();
			let assistrules = this.$refs.assistrules.getData();
			let data = {
				accWeb,
				acc: {
					configuration,
					settings,
					event,
					eventRules,
					entrylist,
					bop,
					assistrules
				},
				accExtraSettings: {
					passwordIsEmpty: settings.passwordIsEmpty,
					adminPasswordIsEmpty: settings.adminPasswordIsEmpty,
					spectatorPasswordIsEmpty: settings.spectatorPasswordIsEmpty
				}
			};
			data.acc.settings.serverName = this.servername;

			let url = "/api/instance";

			if (this.id) {
				url += "/" + this.id;
			}

			axios
				.post(url, data)
				.then(() => {
					this.$router.push("/");
				})
				.catch(e => {
					this.$store.commit("toast", this.$t("save_error"));
				});
		}
	}
};
</script>

<i18n>
{
    "en": {
        "save": "Save",
        "cancel": "Cancel",
        "back": "Back",
        "servername_label": "Servername",
        "server_config": "Configure server",
        "save_error": "Error saving configuration, please check your input."
    }
}
</i18n>
