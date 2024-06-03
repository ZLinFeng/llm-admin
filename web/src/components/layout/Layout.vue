<script setup lang="ts">
import Splitter from "primevue/splitter"
import SplitterPanel from "primevue/splitterpanel"
</script>

<template>
  <Splitter class="app-container" :step="10">
    <SplitterPanel :size="12" :min-size="12">Sider</SplitterPanel>
    <SplitterPanel :size="88" :min-size="88">
      <Splitter layout="vertical">
        <SplitterPanel :size="8" :min-size="8">Header</SplitterPanel>
        <SplitterPanel :size="92" :min-size="92">Main</SplitterPanel>
      </Splitter>
    </SplitterPanel>
  </Splitter>
</template>

<style scoped>
.app-container {
  height: 100vh;
}
</style>