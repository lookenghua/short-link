<template>
  <n-layout has-sider class="layout" position="absolute">
    <LayoutSider />
    <n-layout>
      <n-layout-header> <LayoutHeader /></n-layout-header>
      <n-layout-content content-style="padding: 24px;" class="app-main">
        <router-view v-slot="{ Component }" v-if="appStore.showRouterView">
          <keep-alive :include="includes">
            <component :is="Component" />
          </keep-alive>
        </router-view>
        <Lottie
          v-else
          :animation-data="animationData"
          style="width: 400px; height: 400px; margin: 100px auto 0"
        />
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({ name: "LayoutPage" });
</script>
<script setup>
import { computed } from "vue";
import { useAppStore } from "@/store/modules/app.js";
import { useLayoutStore } from "@/store/modules/layout";
import animationData from "../assets/lottie/40110-waiting-with-a-cup-coffee.json";
import LayoutHeader from "./header/index.vue";
import LayoutSider from "./sider/index.vue";
import Lottie from "@/components/Lottie/Lottie.vue";

const layoutStore = useLayoutStore();
const appStore = useAppStore();
const includes = computed(() =>
  layoutStore.visitedViews.map((it) => it.name).filter((it) => !!it)
);
</script>

<style scoped lang="scss">
.layout {
  .app-main {
    min-height: calc(100vh - 130px);
    transition: all 0.2s;
    background: white;

    ::v-deep(.app-container) {
      padding: 30px 20px;
    }
  }
}
</style>
<style lang="scss">
.right-menu {
  .ant-dropdown-menu-item-icon {
    line-height: 1;
    margin-right: 2px;
    position: relative;
    top: 4px;
  }
}
</style>
