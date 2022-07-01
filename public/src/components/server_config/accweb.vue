<template>
    <collapsible :title="$t('title')">
        <div class="server-settings-container two-columns">
            <div>
                <checkbox :label="$t('autostart_label')" v-model="autoStart"></checkbox>

                <div v-if="os.name != 'windows'">
                    <checkbox :label="$t('enable_adv_windows_conf')" v-model="enableAdvWinCfg"></checkbox>

                    <div v-if="enableAdvWinCfg" style="padding: 10px;">
                        <div class="server-settings-container two-columns">
                            <selection :label="$t('priority_label')" :options="priorities" v-model="priority"></selection>

                            <checkbox :label="$t('enable_windows_firewall')" v-model="enableWindowsFirewall"></checkbox>
                        </div>        
            
                        <label>Core Affinity:</label> <br /> 
                        <div class="server-settings-container four-columns">
                            <checkbox :label="'CPU '+(n-1)" v-for="n in os.numCpu" :key="n" v-model="coreAffinityCPU[n-1]"></checkbox>
                        </div>
                        <div>Empty means ALL</div>
                    </div>
                </div>
            </div>
            
            <div></div>
        </div>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import checkbox from "../checkbox.vue";
import selection from "../selection.vue";

export default {
    components: {collapsible, checkbox, selection},
    data() {
        return {
            enableAdvWinCfg: false,
            enableWindowsFirewall: false,
            autoStart: false,
            priorities: [
                {value: 256, label: "Realtime"},
                {value: 128, label: "High"},
                {value: 32768, label: "Above Normal"},
                {value: 32, label: "Normal"},
                {value: 16384, label: "Below Normal"},
                {value: 64, label: "Low"},
            ],
            priority: 32,
            coreAffinity: 0,
            coreAffinityCPU: [],
            os: {
                name: "",
                numCpu: 0
            }
        };
    },
    methods: {
        hasCPUAffinity(n) {
            return this.coreAffinity & Math.pow(2, n) ? true : false;
        },
        calculatedAffinity() {
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
            this.autoStart = data.accWeb.autoStart;
            this.coreAffinity = data.accWeb.coreAffinity;
            this.os = data.os;

            for (let i = 0; i <= this.os.numCpu; i++) {
                this.coreAffinityCPU[i] = this.hasCPUAffinity(i)
            }
        },
        getData() {
            return {
              autoStart: this.autoStart,
              coreAffinity: this.calculatedAffinity()
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
        "priority_label": "Process priority",
        "enable_windows_firewall": "Enable Windows Firewall"
    }
}
</i18n>
