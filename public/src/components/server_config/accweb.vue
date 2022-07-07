<template>
    <collapsible :title="$t('title')">
        <div class="server-settings-container two-columns">
            <div>
                <checkbox :label="$t('autostart_label')" v-model="autoStart"></checkbox>

                <div v-if="os.name == 'windows'">
                    <checkbox :label="$t('enable_adv_windows_conf')" v-model="enableAdvWindowsCfg"></checkbox>

                    <div v-if="enableAdvWindowsCfg" style="padding: 10px;">
                        <div class="alert">{{$t('adv_windows_alert')}}</div>

                        <div class="server-settings-container two-columns">
                            <selection :label="$t('cpu_priority_label')" :options="priorities" v-model="advWindowsCfg.cpuPriority"></selection>

                            <checkbox :label="$t('enable_windows_firewall')" v-model="advWindowsCfg.enableWindowsFirewall"></checkbox>
                        </div>        
            
                        <label>Core Affinity: (Empty means ALL CPUs)</label> <br /> 
                        <div class="server-settings-container four-columns">
                            <checkbox :label="'CPU '+(n-1)" v-for="n in os.numCpu" :key="n" v-model="coreAffinityCPU[n-1]"></checkbox>
                        </div>
                    </div>
                </div>
            </div>
            
            <div></div>
        </div>
    </collapsible>
</template>

<style>
.alert {
    border: 1px solid #3f0b0b;
    padding: 5px;
    background-color: red;
    font-weight: bold;
    margin-bottom: 10px;
}
</style>

<script>
import collapsible from "../collapsible.vue";
import checkbox from "../checkbox.vue";
import selection from "../selection.vue";
import axios from "axios";

export default {
    components: {collapsible, checkbox, selection},
    data() {
        return {
            autoStart: false,
            enableAdvWindowsCfg: false,
            advWindowsCfg: {
                enableWindowsFirewall: false,
                cpuPriority: 32,
                coreAffinity: 0
            },
            priorities: [
                {value: 256, label: "Realtime"},
                {value: 128, label: "High"},
                {value: 32768, label: "Above Normal"},
                {value: 32, label: "Normal"},
                {value: 16384, label: "Below Normal"},
                {value: 64, label: "Low"},
            ],
            coreAffinityCPU: [],
            os: {
                name: "",
                numCpu: 0
            }
        };
    },
    methods: {
        hasCPUAffinity(n) {
            if (this.advWindowsCfg.coreAffinity === 0) {
                console.log("CPU Affinity was ZERO!");
                this.advWindowsCfg.coreAffinity = Math.pow(2, this.os.numCpu) - 1;
            }

            return this.advWindowsCfg.coreAffinity & Math.pow(2, n) ? true : false;
        },
        calculateAffinity() {
            let total = 0;
            for (const i in this.coreAffinityCPU) {
                if (Object.hasOwnProperty.call(this.coreAffinityCPU, i)) {
                    if (!this.coreAffinityCPU[i]) {
                        continue;
                    }
                    
                    total += Math.pow(2, i);
                }
            }
            return total;
        },
        setData(data) {
            this.autoStart = data.autoStart;
            this.enableAdvWindowsCfg = data.enableAdvWindowsCfg;

            if (data.advWindowsCfg !== null) {
                this.advWindowsCfg = data.advWindowsCfg;
            }

            axios.get("/api/metadata")
                .then(r => {
                    this.os = r.data;

                    for (let i = 0; i <= this.os.numCpu; i++) {
                        this.coreAffinityCPU[i] = this.hasCPUAffinity(i)
                    }
                });
        },
        getData() {
            if (this.enableAdvWindowsCfg) {
                this.advWindowsCfg.coreAffinity = this.calculateAffinity();
                this.advWindowsCfg.cpuPriority = parseInt(this.advWindowsCfg.cpuPriority);
            }

            return {
                autoStart: this.autoStart,
                enableAdvWindowsCfg: this.enableAdvWindowsCfg,
                advWindowsCfg: this.advWindowsCfg
            };
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "ACC Web configuration",
        "autostart_label": "Server instance auto start.",
        "enable_adv_windows_conf": "Advanced Windows Configurations",
        "cpu_priority_label": "Process priority",
        "enable_windows_firewall": "Enable Windows Firewall",
        "adv_windows_alert": "CAUTION: If you are not familiarized with this terms, DISABLE this feature!"
    }
}
</i18n>
