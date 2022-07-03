<template>
<v-container fluid>
			<v-col md="6">
				<v-card>
					<v-list-item three-line>
						<v-list-item-content>
							<div class="text-overline mb-4">
								{{ server.name }}
							<span v-if="!ro">
								<i class="fas fa-terminal" v-on:click="logs" :title="$t('view_logs')"></i>
								<i class="fas fa-copy" v-on:click="copyConfig" v-if="is_admin" :title="$t('copy_config')"></i>
								<i class="fas fa-file-download" v-on:click="exportConfig" :title="$t('export_config')"></i>
							</span>
							<v-divider></v-divider>
							</div>
							<v-list-item-title class="text-h5 mb-1" v-if="server.pid">
								PID: {{ server.pid }}
								<v-divider></v-divider>
							</v-list-item-title>
							<v-list-item-title class="text-h6 mb-4">
								UDP: {{ server.udpPort }} TCP: {{ server.tcpPort }}
								<v-divider></v-divider>
								{{ $t("track") }}: {{ server.track }}
							</v-list-item-title>
							<v-list-item-title class="text-h5 mb-1" v-if="!ro">
							<v-divider></v-divider>
								{{ $t("configuration_directory") }}: {{ server.id }}
							</v-list-item-title>
							<div class="info state" v-if="server.pid">
							<v-divider></v-divider>
								<b>{{ $t("state") }}: </b>{{ $t(server.serverState) }} &bull;
								<b>{{ $t("number_of_drivers") }}: </b
								>{{ formattedServerClientCount }} &bull;
								<b>{{ $t("session") }}: </b>
								<span v-if="server.sessionType"
									>{{ server.sessionType }} ({{ server.sessionPhase }}) -
									{{ server.sessionRemaining }} min(s)</span
								>
								<span v-else>{{ $t("not_detected") }}</span>
							</div>
						</v-list-item-content>
					</v-list-item>

					<v-card-actions>
						<v-btn
							class="start"
							v-on:click="start"
							v-if="is_mod && !ro && !server.pid"
							>{{ $t("start_server") }}</v-btn
						>
						<v-btn
							class="stop"
							v-on:click="stop"
							v-if="is_mod && !ro && server.pid"
							>{{ $t("stop_server") }}</v-btn
						>
						<v-btn v-on:click="deleteServer" v-if="is_admin && !ro"><i class="fas fa-trash"></i>{{ $t('delete_server')}}</v-btn>
						<v-btn v-on:click="edit" v-if="is_admin && !ro"><i class="fas fa-cog"></i>{{ $t('change_config')}}</v-btn>						
						<div class="online" v-if="ro && server.pid">{{ $t("running") }}</div>
						<div class="offline" v-if="ro && !server.pid">{{ $t("offline") }}</div>
					</v-card-actions>
				</v-card>
			</v-col>
</v-container>		
</template>

<style>
.state {
	margin-top: 10px;
}

.state b {
	color: #00ff14;
}
.container {display: flex}
</style>

<script>
import axios from "axios";

export default {
	props: ["server", "ro"],
	data: () => ({
		justify: ["start", "center", "end", "space-around", "space-between"],
        value: 1
	}),
	computed: {
		formattedServerClientCount: function() {
			return this.server.serverState === "not_registered"
				? "-"
				: this.server.nrClients;
		}
	},
	methods: {
		edit() {
			this.$router.push(`/server?id=${this.server.id}`);
		},
		logs() {
			this.$router.push(`/logs?id=${this.server.id}`);
		},
		live() {
			this.$router.push(`/live?id=${this.server.id}`);
		},
		copyConfig() {
			axios
				.post(`/api/instance/${this.server.id}/clone`)
				.then(() => {
					this.$emit("copied");
				})
				.catch(e => {
					this.$store.commit("toast", this.$t("copy_server_error"));
				});
		},
		exportConfig() {
			let link = document.createElement("a");
			link.setAttribute("type", "hidden");
			link.href = `/api/instance/${this.server.id}/export?token=${
				this.$store.state.auth.token
			}`;
			document.body.appendChild(link);
			link.click();
			link.remove();
		},
		deleteServer() {
			axios
				.delete(`/api/instance/${this.server.id}`)
				.then(() => {
					this.$emit("deleted");
				})
				.catch(e => {
					this.$store.commit("toast", this.$t("delete_server_error"));
				});
		},
		start() {
			axios
				.post(`/api/instance/${this.server.id}/start`)
				.then(() => {
					this.$emit("started");
				})
				.catch(e => {
					this.$store.commit("toast", this.$t("start_server_error"));
				});
		},
		stop() {
			axios
				.post(`/api/instance/${this.server.id}/stop`)
				.then(() => {
					this.$emit("stopped");
				})
				.catch(e => {
					this.$store.commit("toast", this.$t("stop_server_error"));
				});
		}
	}
};
</script>

<i18n>
{
    "en": {
        "start_server": "Start",
        "stop_server": "Stop",
        "change_config": "Change config",
        "view_logs": "View logs",
        "view_live": "View live",
        "copy_config": "Copy config",
        "export_config": "Export config",
        "delete_server": "Delete server",
        "copy_server_error": "Error copying server configuration.",
        "delete_server_error": "Error deleting server configuration.",
        "start_server_error": "Error starting server, please check the logs.",
        "stop_server_error": "Error stopping server.",
        "track": "Track",
        "configuration_directory": "Config dir",
        "running": "Running",

        "state": "State",
        "number_of_drivers": "Drivers",
        "session": "Session",
        "not_detected": "Not detected",

        "offline": "Offline",
        "starting": "Starting",
        "not_registered": "Waiting for events",
        "online": "Online"
    }
}
</i18n>
