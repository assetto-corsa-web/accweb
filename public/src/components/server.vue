<template>
			<v-col md="6">			
				<v-card>
					<div data-app>
					          <v-app-bar dense
            flat
            color="rgba(0, 0, 0, 0)"
          >

            <v-toolbar-title class="text-h6 white--text pl-0 text-center">
              ID : {{ server.id }}
            </v-toolbar-title>

            <v-spacer></v-spacer>

      <v-menu
        bottom
        left
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            color="yellow"
            v-bind="attrs"
            v-on="on"
          >
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item>
								<v-btn
									elevation="2"
									x-small
									color="primary"
									rounded	
									v-on:click="deleteServer"
									v-if="is_admin && !ro"
									><i class="fas fa-trash"></i>Del</v-btn
								>
          </v-list-item>

          <v-list-item>
								<v-btn
									elevation="2"
									x-small
									color="primary"
									rounded
									v-on:click="edit"
									v-if="is_admin && !ro"
									><i class="fas fa-cog"></i>Settings</v-btn
								>
          </v-list-item>

          <v-list-item>
								<v-btn
									elevation="2"
									x-small
									color="primary"
									rounded
									v-on:click="logs"
									v-if="is_admin && !ro"
									><i class="fas fa-terminal"></i>Logs</v-btn
								>	
          </v-list-item>

          <v-list-item>
								<v-btn
									elevation="2"
									x-small
									color="primary"
									rounded						
									v-on:click="copyConfig"
									v-if="is_admin && !ro"
									><i class="fas fa-copy"></i>Copy CFG</v-btn
								>
          </v-list-item>

          <v-list-item>
            								<v-btn
									elevation="2"
									x-small
									color="primary"
									rounded
									v-on:click="exportConfig"
									v-if="is_admin && !ro"
									><i class="fas fa-file-download"></i>Download CFG</v-btn
								>
          </v-list-item>
        </v-list>
      </v-menu>
          </v-app-bar>
					<v-list-item active-class>
						<v-list-item-content>
							<div class="text-overline mb-4">
								{{ server.name }}
								<v-divider></v-divider>
							</div>
							<v-list-item-title class="text-h6 mb-1" v-if="server.pid">
								PID: {{ server.pid }}
								<v-divider></v-divider>
							</v-list-item-title>
							<v-list-item-title class="text-h6 mb-1">
								UDP: {{ server.udpPort }} TCP: {{ server.tcpPort }}
								<v-divider></v-divider>
								{{ $t("track") }}: {{ server.track }}
								<v-divider></v-divider>
							</v-list-item-title>
							<div class="info state" v-if="server.pid">
								<b>{{ $t("state") }}: </b>{{ $t(server.serverState) }} &bull;
								<b>{{ $t("number_of_drivers") }}: </b>
								{{ formattedServerClientCount }}
								<v-divider></v-divider>
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
						<v-spacer></v-spacer>
						<v-row dense>
							<v-btn
								elevation="2"
								x-small
								color="primary"
								rounded
								class="start"
								v-on:click="start"
								v-if="is_mod && !ro && !server.pid"
								><i class="fas fa-play"></i>{{ $t("start_server") }}</v-btn
							>
								<v-btn
									elevation="2"
									x-small
									color="primary"
									rounded									
									class="stop"
									v-on:click="stop"
									v-if="is_mod && !ro && server.pid"
									><i class="fas fa-stop"></i>{{ $t("stop_server") }}</v-btn
								>

								<v-btn
									elevation="2"
									x-small
									color="primary"
									rounded
									v-on:click="live"
									v-if="is_mod && !ro && server.pid"
									><i class="fas fa-tv"></i>{{ $t("view_live") }}</v-btn
								>

							<div class="online" v-if="ro && server.pid">{{ $t("running") }}</div>
							<div class="offline" v-if="ro && !server.pid">{{ $t("offline") }}</div>
			
						</v-row>
					</v-card-actions>
					</div>
				</v-card>
			</v-col>
</template>

<style>
.state {
	margin-top: 10px;
}

.state b {
	color: #00ff14;
}

input, select, textarea {
    background-color: #ffffff00;
    border-style: solid;
}

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
