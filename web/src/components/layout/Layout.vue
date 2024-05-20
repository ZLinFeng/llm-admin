<script setup lang="ts">

import Aside from "@/components/layout/Aside.vue"
import Content from "@/components/layout/Content.vue"
</script>

<template>
  <div>
    <el-container>
      <el-aside class="aside-wrapper">
        <Aside />
      </el-aside>
      <el-container>
        <el-header>
          <Header />
        </el-header>
        <el-main>
          <Content />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
.aside-wrapper {
  height: 100vh;
  width: 200px;
}
</style>