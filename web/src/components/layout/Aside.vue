<script setup lang="ts">

import {Location, User} from "@element-plus/icons-vue"
</script>

<template>
  <div class="aside-wrapper">
    <div class="logo-wrapper">
      <div>Logo</div>
    </div>
    <div class="menu-wrapper">
      <el-scrollbar>
        <el-menu
          active-text-color="#ffd04b"
          background-color="#545c64"
          text-color="#fff"
        >
          <el-sub-menu index="1">
            <template #title>
              <el-icon><Location /></el-icon>
              <span>Location</span>
            </template>
            <el-menu-item index="1-1">
              item one
            </el-menu-item>
            <el-menu-item index="1-2">
              item two
            </el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="2">
            <template #title>
              <el-icon><User /></el-icon>
              <span>User</span>
            </template>
            <el-menu-item index="2-1">
              item one
            </el-menu-item>
            <el-menu-item index="2-2">
              item two
            </el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title>
              <el-icon><User /></el-icon>
              <span>User</span>
            </template>
            <el-menu-item index="3-1">
              item one
            </el-menu-item>
            <el-menu-item index="3-2">
              item two
            </el-menu-item>
            <el-menu-item index="3-1">
              item one
            </el-menu-item>
            <el-menu-item index="3-2">
              item two
            </el-menu-item>
            <el-menu-item index="3-1">
              item one
            </el-menu-item>
            <el-menu-item index="3-2">
              item two
            </el-menu-item>
            <el-menu-item index="3-1">
              item one
            </el-menu-item>
            <el-menu-item index="3-2">
              item two
            </el-menu-item>
            <el-menu-item index="3-1">
              item one
            </el-menu-item>
            <el-menu-item index="3-2">
              item two
            </el-menu-item>
            <el-menu-item index="3-1">
              item one
            </el-menu-item>
            <el-menu-item index="3-2">
              item two
            </el-menu-item>
            <el-menu-item index="3-1">
              item one
            </el-menu-item>
            <el-menu-item index="3-2">
              item two
            </el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-scrollbar>
    </div>
    <div class="logout-wrapper">
      Logout
    </div>
  </div>
</template>

<style scoped>
.aside-wrapper {
  height: 100vh;
}

.logo-wrapper {
  height: 10%;
  background-color: burlywood;
}

.logout-wrapper {
  height: 5%;
  background-color: burlywood;
}

.menu-wrapper {
  height: 85%;
  background-color: #545c64;
}
</style>