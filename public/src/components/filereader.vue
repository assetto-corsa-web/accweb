<template>
  <label class="text-reader">
    <i class="fas fa-file-upload" :title="$t('upload_config')"></i>
    <input type="file" @change="loadTextFromFile">
  </label>
</template>

<style>
  input[type="file"] {
    display: none;
  }

  .text-reader {
    width: auto;
    cursor: pointer;
  }
</style>

<script>
export default {
  methods: {
    loadTextFromFile(ev) {
      const file = ev.target.files[0];
      const reader = new FileReader();

      reader.onload = e => {
          this.$emit("load", e.target.result);
          ev.target.value = null;
      }
      reader.readAsText(file);
    }
  }
};
</script>

<i18n>
{
  "en": {
    "upload_config": "Upload config"
  }
}
</i18n>
