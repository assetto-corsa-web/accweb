<template>
<layout small="true">
    <div class="title">
        <h2>{{ $t("title") }}</h2>
    </div>
    <p>{{ $t("message") }}</p>
</layout>
</template>

<script>
import {
    layout,
    end
} from "../components";

export default {
    components: {
        layout,
        end
    }
};
</script>

<i18n>
{

"en": {

"title": "#404",

"message": "Page not found."

}
}
</i18n>
