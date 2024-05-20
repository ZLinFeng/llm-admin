<script setup lang="ts">

import {DArrowLeft, DArrowRight, Location, User} from "@element-plus/icons-vue"
import {ref} from "vue"

const collapse = ref(false)
</script>

<template>
  <div class="aside-wrapper">
    <div class="logo-wrapper">
      <span>LLM-Admin</span>
    </div>
    <div class="menu-wrapper">
      <el-scrollbar>
        <el-menu
          active-text-color="#ffd04b"
          background-color="#545c64"
          text-color="#fff"
          class="menu"
          :collapse="collapse"
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
      <div
        class="collapse-btn"
        @click="collapse = !collapse"
      >
        <el-icon v-if="!collapse">
          <DArrowLeft />
        </el-icon>
        <el-icon v-else>
          <DArrowRight />
        </el-icon>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.aside-wrapper {
  height: 100vh;
}

.logo-wrapper {
  height: 10%;
  background-color: #545c64;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #fff;
  font-size: 1.5em;
}

.logout-wrapper {
  height: 5%;
  background-color: #545c64;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 15px;
  color: white;
  cursor: pointer;
}

.collapse-btn {
  width: 40px;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 20px;
}

.collapse-btn:hover {
  background-color: #ffd04b;
  color: black;
}

.menu-wrapper {
  height: 85%;
  background-color: #545c64;
}

.menu {
  border: 0;
}

.el-sub-menu__title {
  span {
    font-size: 1.2em;
  }
}

.el-menu-item {
  font-size: 1em;
}
</style>