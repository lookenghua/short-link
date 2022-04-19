<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <router-view />
  </n-config-provider>
</template>

<script setup>
import { zhCN, dateZhCN } from "naive-ui";
</script>

<style></style>
