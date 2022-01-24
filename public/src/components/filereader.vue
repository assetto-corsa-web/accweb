<template>
  <label class="text-reader">
    <input type="file" @change="loadTextFromFile">
  </label>
</template>

<style>
.text-reader {
  position: relative;
  overflow: hidden;
  display: inline-block;

  /* Fancy button looking */
  border: 2px solid black;
  border-radius: 5px;
  padding: 8px 12px;
  cursor: pointer;
  width: 100px;
  background-color: #c7d5e0;
}
.text-reader input {
  position: absolute;
  top: 0;
  left: 0;
  z-index: -1;
  opacity: 0;
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
