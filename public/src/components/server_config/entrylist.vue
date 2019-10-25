<template>
    <collapsible :title="$t('title')">
        <entry v-for="entry in entries"
            :key="entry.index"
            :entry="entry"
            v-on:remove="removeEntry"></entry>
        <field type="number" :label="$t('forceentrylist_label')" v-model="forceEntryList"></field>
        <field type="number" :label="$t('defaultgridposition_label')" v-model="defaultGridPosition"></field>
        <button v-on:click="addEntry">{{$t("add_entry_button")}}</button>
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";
import entry from "./entry.vue";
import field from "../field.vue";

export default {
    components: {collapsible, entry, field},
    data() {
        return {
            entryIndex: 0,
            entries: [],
            forceEntryList: 0,
            defaultGridPosition: 0
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
                    overrideCarModelForCustomCar: entries[i].overrideCarModelForCustomCar
                });
                this.entryIndex++;
            }

            this.forceEntryList = data.forceEntryList;
            this.defaultGridPosition = data.defaultGridPosition;
        },
        getData() {
            return {
                entries: this.getEntryData(),
                forceEntryList: parseInt(this.forceEntryList),
                defaultGridPosition: parseInt(this.defaultGridPosition)
            };
        },
        getEntryData() {
            let entries = [];

            for(let i = 0; i < this.entries.length; i++) {
                entries.push({
                    drivers: this.entries[i].drivers,
                    raceNumber: parseInt(this.entries[i].raceNumber),
                    forcedCarModel: parseInt(this.entries[i].forcedCarModel),
                    overrideDriverInfo: parseInt(this.entries[i].overrideDriverInfo),
                    isServerAdmin: parseInt(this.entries[i].isServerAdmin),
                    customCar: this.entries[i].customCar,
                    overrideCarModelForCustomCar: parseInt(this.entries[i].overrideCarModelForCustomCar)
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
                overrideDriverInfo: 0,
                isServerAdmin: 0,
                customCar: "",
                overrideCarModelForCustomCar: 0
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
        }
    }
}
</script>

<i18n>
{
    "en": {
        "title": "Entry List",
        "forceentrylist_label": "Force Entry List",
        "defaultgridposition_label": "Default Grid Position",
        "add_entry_button": "Add Entry"
    }
}
</i18n>
