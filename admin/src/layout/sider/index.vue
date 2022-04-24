<template>
  <n-layout-sider
    collapse-mode="width"
    :collapsed-width="48"
    :width="200"
    :collapsed="collapsed"
    @collapse="collapsed = true"
    @expand="collapsed = false"
    class="sider"
    :inverted="true"
  >
    <div class="logo-area flex">
      <img src="../../assets/images/logo.png" alt="" class="logo" />
      <span class="title">博客管理后台</span>
    </div>
    <n-menu
      :collapsed="collapsed"
      :collapsed-width="64"
      :collapsed-icon-size="22"
      :options="menus"
      :value="openKeys"
      @update:expanded-keys="handleUpdateExpandedKeys"
      @update:value="handleMenuClick"
    />
  </n-layout-sider>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({ name: "LayoutSider" });
</script>
<script setup>
import { computed, ref, toRaw } from "vue";
import { useLayoutStore } from "@/store/modules/layout";
import { usePermissionStore } from "@/store/modules/perssion";
import { generatorMenu } from "@/utils";
const layoutStore = useLayoutStore();
const permissionStore = usePermissionStore();
const menus = ref([]);

let openKeys = computed({
  get() {
    return layoutStore.openKeys;
  },
  set(val) {
    layoutStore.openKeys = val;
  },
});
// const selectedKeys = computed({
//   get() {
//     return layoutStore.activeMenu;
//   },
//   set(val) {
//     layoutStore.setActiveMenu(val);
//   },
// });

const collapsed = computed({
  get() {
    return layoutStore.collapsed;
  },
  set(isCollapsed) {
    layoutStore.collapsed = isCollapsed;
  },
});
const routes = permissionStore.permissionRoutes.sort(
  (a, b) => a.meta.sort - b.meta.sort
);

menus.value = generatorMenu(toRaw(routes));

// 点击菜单
function handleMenuClick(key) {
  console.log(`当前点击菜单: ${key}`);
}

// 展开菜单
function handleUpdateExpandedKeys(keys) {
  console.log(keys);
}
</script>
<style scoped lang="scss">
.sider {
  height: 100vh;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  .logo-area {
    padding: 10px 5px;
    white-space: nowrap;
    overflow: hidden;
    width: 200px;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

    .logo {
      width: 38px;
      transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
      margin-right: 5px;
    }

    .title {
      display: inline-block;
      font-size: 20px;
      padding-top: 5px;
      opacity: 1;
      transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }
  }
  ::v-deep(.n-menu-item-content__icon) {
    color: white;
  }
}
</style>
