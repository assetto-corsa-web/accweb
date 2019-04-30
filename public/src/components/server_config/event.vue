<template>
    <collapsible :title="$t('title')">
        TODO
    </collapsible>
</template>

<script>
import collapsible from "../collapsible.vue";

export default {
    components: {collapsible}
}
</script>

<i18n>
{
    "en": {
        "title": "Event settings"
    }
}
</i18n>
