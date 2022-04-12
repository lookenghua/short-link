<template>
  <a-layout-sider
    v-model:collapsed="collapsed"
    :collapsedWidth="48"
    :trigger="null"
    collapsible
    class="sider"
  >
    <div class="logo-area flex">
      <img src="../../assets/images/logo.png" alt="" class="logo" />
      <span class="title">博客管理后台</span>
    </div>
    <a-menu
      v-model:selectedKeys="selectedKeys"
      v-model:openKeys="openKeys"
      theme="dark"
      mode="inline"
      @click="handleMenuClick"
    >
      <MenuItem v-for="item in routes" :key="item.path" :route="item" />
    </a-menu>
  </a-layout-sider>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({ name: "LayoutSider" });
</script>
<script setup>
import { computed } from "vue";
import { useLayoutStore } from "@/store/modules/layout";
import { usePermissionStore } from "@/store/modules/perssion";
import MenuItem from "./components/MenuItem.vue";

const layoutStore = useLayoutStore();
const permissionStore = usePermissionStore();
let openKeys = computed({
  get() {
    return layoutStore.collapsed ? [] : layoutStore.openKeys;
  },
  set(val) {
    layoutStore.openKeys = val;
  },
});
const selectedKeys = computed({
  get() {
    return layoutStore.activeMenu;
  },
  set(val) {
    layoutStore.setActiveMenu(val);
  },
});

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

// 点击菜单
function handleMenuClick({ key }) {
  console.log(`当前点击菜单: ${key}`);
}
</script>

<style scoped lang="scss">
.sider {
  height: 100vh;
  color: white;
  position: fixed;
  left: 0;

  .logo-area {
    color: white;
    padding: 10px 5px;
    white-space: nowrap;
    overflow: hidden;

    .logo {
      width: 38px;
    }

    .title {
      display: inline-block;
      margin-left: 5px;
      font-size: 20px;
      padding-top: 5px;
      opacity: 1;
      transition: opacity 0.3s cubic-bezier(0.645, 0.045, 0.355, 1), margin 0.3s,
        color 0.3s;
    }
  }

  .ant-layout-sider-collapsed {
    --icon-size: 40px;
  }

  .ant-menu-inline-collapsed {
    ::v-deep(.ant-menu-item) {
      display: flex;
      //line-height: 48px;
    }
  }
}
</style>
<style lang="scss">
.ant-menu-submenu-popup {
  .ant-menu-item {
    display: flex;
  }
}
</style>
