<template>
  <nConfig-provider
    :locale="zhCN"
    :date-locale="dateZhCN"
    :theme-overrides="themeOverrides"
  >
    <n-message-provider> <router-view /> </n-message-provider>
  </nConfig-provider>
</template>

<script setup>
import { zhCN, dateZhCN } from "naive-ui";
/**
 * js 文件下使用这个做类型提示
 * @type import('naive-ui').GlobalThemeOverrides
 */
const themeOverrides = {
  common: {},
  Menu: {
    dividerColor: "rgb(239, 239, 245)",
    bezier: "cubic-bezier(0.4, 0, 0.2, 1)",
    fontSize: "14px",
    borderColorHorizontal: "#000",
    borderRadius: "3px",
    itemHeight: "42px",
    groupTextColor: "#AAA",
    color: "#000",
    itemTextColor: "#BBB",
    itemTextColorHover: "#FFF",
    itemTextColorActive: "#FFF",
    itemTextColorChildActive: "#FFF",
    itemTextColorActiveHover: "#FFF",
    itemIconColor: "#BBB",
    itemIconColorHover: "#FFF",
    itemIconColorActive: "#FFF",
    itemIconColorActiveHover: "#FFF",
    itemIconColorChildActive: "#FFF",
    itemIconColorCollapsed: "#BBB",
    itemTextColorHorizontal: "#BBB",
    itemTextColorHoverHorizontal: "#FFF",
    itemTextColorActiveHorizontal: "#FFF",
    itemTextColorChildActiveHorizontal: "#FFF",
    itemTextColorActiveHoverHorizontal: "#FFF",
    itemIconColorHorizontal: "#BBB",
    itemIconColorHoverHorizontal: "#FFF",
    itemIconColorActiveHorizontal: "#FFF",
    itemIconColorActiveHoverHorizontal: "#FFF",
    itemIconColorChildActiveHorizontal: "#FFF",
    arrowColor: "#BBB",
    arrowColorHover: "#FFF",
    arrowColorActive: "#FFF",
    arrowColorActiveHover: "#FFF",
    arrowColorChildActive: "#FFF",
    itemColorHover: "#000",
    itemColorActive: "#2d8cf0",
    itemColorActiveHover: "#2d8cf0",
    itemColorActiveCollapsed: "#2d8cf0",
  },
};
</script>

<style></style>
