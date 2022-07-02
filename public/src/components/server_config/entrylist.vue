<template>
    <collapsible :title="$t('title')" with-import="true" import-filename="entrylist.json" @load="setData">
        <entry v-for="entry in entries"
            :key="entry.index"
            :entry="entry"
            v-on:remove="removeEntry"></entry>
        <checkbox :label="$t('forceentrylist_label')" v-model="forceEntryList"></checkbox>
        <button v-on:click="addEntry">{{$t("add_entry_button")}}</button>
        <button v-on:click="clearEntries">{{$t("clear_entries_button")}}</button>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import entry from "./entry.vue";
import field from "../field.vue";
import checkbox from "../checkbox.vue";

export default {
    components: {collapsible, entry, field, checkbox},
    data() {
        return {
            entryIndex: 0,
            entries: [],
            forceEntryList: 0
        };
    },
    methods: {
        setData(data) {
            let entries = data.entries;

            for(let i = 0; i < entries.length; i++) {
                this.entries.push({
                  index: this.entryIndex,
                  drivers: entries[i].drivers,
                  raceNumber: entries[i].raceNumber,
                  forcedCarModel: entries[i].forcedCarModel,
                  overrideDriverInfo: entries[i].overrideDriverInfo,
                  isServerAdmin: entries[i].isServerAdmin,
                  customCar: entries[i].customCar,
                  overrideCarModelForCustomCar: entries[i].overrideCarModelForCustomCar,
                  ballastKg: entries[i].ballastKg,
                  restrictor: entries[i].restrictor,
                  defaultGridPosition: entries[i].defaultGridPosition
                });
                this.entryIndex++;
            }
            this.forceEntryList = data.forceEntryList;
        },
        getData() {
            return {
                entries: this.getEntryData(),
                forceEntryList: this.forceEntryList ? 1 : 0
            };
        },
        getEntryData() {
            let entries = [];

            for(let i = 0; i < this.entries.length; i++) {
                entries.push({
                  drivers: this.entries[i].drivers,
                  raceNumber: parseInt(this.entries[i].raceNumber),
                  forcedCarModel: parseInt(this.entries[i].forcedCarModel),
                  overrideDriverInfo: this.entries[i].overrideDriverInfo ? 1 : 0,
                  isServerAdmin: this.entries[i].isServerAdmin ? 1 : 0,
                  customCar: this.entries[i].customCar,
                  overrideCarModelForCustomCar: this.entries[i].overrideCarModelForCustomCar ? 1 : 0,
                  ballastKg: parseInt(this.entries[i].ballastKg),
                  restrictor: parseInt(this.entries[i].restrictor),
                  defaultGridPosition: parseInt(this.entries[i].defaultGridPosition)
                });
            }

            return entries;
        },
        addEntry() {
            this.entries.push({
              index: this.entryIndex,
              drivers: [],
              raceNumber: 0,
              forcedCarModel: -1,
              overrideDriverInfo: false,
              isServerAdmin: false,
              customCar: "",
              overrideCarModelForCustomCar: false,
              ballastKg: 0,
              restrictor: 0,
              defaultGridPosition: -1
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
        "title": "Entry List",
        "forceentrylist_label": "Force Entry List",
        "add_entry_button": "Add Entry",
        "clear_entries_button": "Clear all Entries",
        "confirm_clear_entries": "Do you really want to remove all entries?"
    }
}
</i18n>
