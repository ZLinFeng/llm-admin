<script setup lang="ts">

import Aside from "@/components/layout/Aside.vue"
import Content from "@/components/layout/Content.vue"
</script>

<template>
  <div>
    <el-container>
      <el-aside class="aside-container">
        <Aside />
      </el-aside>
      <el-container>
        <el-header class="header-container">
          <Header />
        </el-header>
        <el-main class="main-container">
          <Content />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
.aside-container {
  height: 100vh;
}

.header-container {
  padding: 0;
}
</style>