<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider> <router-view /> </n-message-provider>>
  </n-config-provider>
</template>

<script setup>
import { zhCN, dateZhCN, useMessage } from "naive-ui";
</script>

<style></style>
