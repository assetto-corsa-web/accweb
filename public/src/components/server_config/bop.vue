<template>
    <collapsible :title="$t('title')" with-import="true" import-filename="bop.json" @load="setData">
        <entry v-for="entry in entries"
            :key="entry.index"
            :entry="entry"
            v-on:remove="removeEntry"></entry>
        <button v-on:click="addEntry">{{$t("add_entry_button")}}</button>
        <button v-on:click="clearEntries">{{$t("clear_entries_button")}}</button>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import field from "../field.vue";
import selection from "../selection.vue";
import entry from "./bopsession.vue";

export default {
    components: {collapsible, field, selection, entry},
    data() {
    	return {
    		entryIndex: 0,
            configVersion: 1,
            entries: []
    	};
    },
    methods: {
        setData(data) {
            this.track = data.track;
            this.carModel = data.carModel;
            this.ballastKg = data.ballastKg;
            this.restrictor = data.restrictor;
            this.setEntryData(data.entries);
        },
        setEntryData(data) {
            for(let i = 0; i < data.length; i++) {
                this.entries.push({
                    index: this.entryIndex,
                    track: data[i].track,
                    carModel: data[i].carModel,
                    ballastKg: data[i].ballastKg,
                    restrictor: data[i].restrictor                    
                });
                this.entryIndex++;
            }
        },
    	getData() {
    		return {
				track: this.track,
				carModel: parseInt(this.carModel),
				ballastKg: parseInt(this.ballastKg),
				restrictor: parseInt(this.restrictor),				
                entries: this.getEntryData(),
                configVersion: 1
    		};
    	},
        getEntryData() {
            let entries = [];

            for(let i = 0; i < this.entries.length; i++) {
                entries.push({
                    track: this.entries[i].track,
                    carModel: parseInt(this.entries[i].carModel),
                    ballastKg: parseInt(this.entries[i].ballastKg),
					restrictor: parseInt(this.entries[i].restrictor)
                });
            }

            return entries;
        },
        addEntry() {
            this.entries.push({
                index: this.entryIndex,
                track: "",
                carModel: 99,
                ballastKg: 0,
                restrictor: 0
            });
            this.entryIndex++;
        },
        removeEntry(index) {
            index = parseInt(index);

            for(let i = 0; i < this.entries.length; i++) {
                if(this.entries[i].index === index) {
                    this.entries.splice(i, 1);
                    break;
                }
            }
        },
        toFloat(value) {
            if(typeof value === "string") {
                return parseFloat(value.replace(",", "."));
            }

            return value;
        },
        clearEntries(){
            if (!window.confirm(this.$t("confirm_clear_entries"))) {
                return;
            }
            while(this.entries.length > 0) {
                this.entries.splice(this.entries[this.entries.length - 1], 1);
            }
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "BOP Settings",
        "track_label": "Track",
        "carModel_label": "Car Model #",
        "ballast_label": "Ballast: 0 to 100kg max.",
        "restrictor_label": "Motor Restrictor: 0 to 20% max.",
        "add_entry_button": "Add BOP",
        "clear_entries_button": "Clear all BOP",
        "confirm_clear_entries": "Do you really want to remove all BOP?"    
    }
}
</i18n>
