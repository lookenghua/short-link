<template>
  <a-layout class="layout">
    <LayoutSider />
    <a-layout>
      <LayoutHeader />
      <a-layout-content class="app-main">
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
      </a-layout-content>
    </a-layout>
  </a-layout>
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
const leftDistance = computed(() => (layoutStore.collapsed ? "49px" : "201px"));
</script>

<style scoped lang="scss">
.layout {
  .app-main {
    min-height: calc(100vh - 130px);
    margin: 81px 0 0;
    margin-left: v-bind(leftDistance);
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
