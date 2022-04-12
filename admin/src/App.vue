<template>
  <AConfigprovider :locale="zhCN">
    <router-view />
  </AConfigprovider>
</template>

<script setup>
import { ConfigProvider as AConfigprovider } from "ant-design-vue";
import zhCN from "ant-design-vue/es/locale/zh_CN";
</script>

<style></style>
