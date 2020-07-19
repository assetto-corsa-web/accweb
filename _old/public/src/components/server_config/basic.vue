<template>
    <collapsible :title="$t('title')">
        <field type="number" :label="$t('udp_label')" v-model="udpPort"></field>
        <field type="number" :label="$t('tcp_label')" v-model="tcpPort"></field>
        <field type="number" :label="$t('maxconnections_label')" v-model="maxConnections"></field>
        <field type="number" :label="$t('registertolobby_label')" v-model="registerToLobby"></field>
        <field type="number" :label="$t('landiscovery_label')" v-model="lanDiscovery"></field>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";

export default {
    components: {collapsible, field},
    data() {
        return {
            udpPort: 9600,
            tcpPort: 9600,
            maxConnections: 10,
            registerToLobby: 1,
            lanDiscovery: 0
        };
    },
    methods: {
        setData(data) {
            this.udpPort = data.udpPort;
            this.tcpPort = data.tcpPort;
            this.maxConnections = data.maxConnections;
            this.registerToLobby = data.registerToLobby;
            this.lanDiscovery = data.lanDiscovery;
        },
        getData() {
            return {
                udpPort: parseInt(this.udpPort),
                tcpPort: parseInt(this.tcpPort),
                maxConnections: parseInt(this.maxConnections),
                registerToLobby: parseInt(this.registerToLobby),
                lanDiscovery: parseInt(this.lanDiscovery)
            };
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Basic configuration",
        "udp_label": "UDP port",
        "tcp_label": "TCP port",
        "maxconnections_label": "Max. connections",
        "registertolobby_label": "Register to lobby",
        "landiscovery_label": "LAN Discovery"
    }
}
</i18n>
