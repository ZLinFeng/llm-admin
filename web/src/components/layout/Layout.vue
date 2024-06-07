<script setup lang="ts">
import Splitter from "primevue/splitter"
import SplitterPanel from "primevue/splitterpanel"
</script>

<template>
  <Splitter class="app-container" :gutter-size="0" layout="vertical">
    <SplitterPanel :size="8" :min-size="8">Header</SplitterPanel>
    <SplitterPanel :size="92" :min-size="92">Main</SplitterPanel>
  </Splitter>
</template>

<style scoped>
.app-container {
  height: 100vh;
}
</style>