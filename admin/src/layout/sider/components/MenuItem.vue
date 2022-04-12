<template>
  <template v-if="hasChildren(props.route)">
    <a-sub-menu
      v-if="children.length > 1 || props.route.alwaysShow"
      :key="props.route.path"
    >
      <template #icon>
        <div v-if="props.route.meta?.icon" :style="menuIconStyle">
          <Icon :icon="props.route.meta?.icon" :style="iconStyle" />
        </div>
      </template>
      <template #title>
        {{ props.route.meta?.title }}
      </template>
      <MenuItem
        v-for="item in children"
        :key="item.name"
        :route="item"
        :parent-path="currentPath"
      />
    </a-sub-menu>
    <a-menu-item
      v-else
      :key="children[0].path"
      @click="handleMenuItemClick(children[0])"
    >
      <template #icon>
        <div v-if="children[0].meta?.icon" :style="menuIconStyle">
          <Icon :icon="children[0].meta?.icon" :style="iconStyle" />
        </div>
      </template>
      {{ children[0].meta?.title }}
    </a-menu-item>
  </template>
  <a-menu-item
    v-else
    :key="props.route.path"
    @click="handleMenuItemClick(props.route)"
  >
    <template #icon>
      <div v-if="props.route.meta?.icon" :style="menuIconStyle">
        <Icon :icon="props.route.meta?.icon" :style="iconStyle" />
      </div>
    </template>
    {{ props.route.meta?.title }}
  </a-menu-item>
</template>

<script setup>
import { reactive } from "vue";
import { useRouter } from "vue-router";
import MenuItem from "../components/MenuItem.vue";

const props = defineProps({
  parentPath: {
    type: String,
    default: "",
  },
  route: {
    type: Object,
    default: () => ({}),
  },
});
const router = useRouter();
// eslint-disable-next-line vue/no-setup-props-destructure
const currentPath = props.route.path;
const children = props.route.children
  ? props.route.children
      .filter((it) => !it.meta?.hidden)
      .sort((a, b) => a.meta.sort - b.meta.sort)
      .map((it) => ({
        ...it,
        path: getFullPath(currentPath, it),
      }))
  : [];

const iconStyle = reactive({
  fontSize: 22,
});

const menuIconStyle = reactive({
  display: "flex",
  alignItems: "center",
  height: "40px",
  minWidth: "20px",
});

// 获取路由全路径
function getFullPath(parentPath, route) {
  if (route.path.startsWith("/")) {
    return route.path;
  } else {
    if (parentPath === "/") {
      if (route.path === "") {
        return ``;
      } else {
        return `/${route.path}`;
      }
    } else {
      if (route.path === "") {
        return `${parentPath}`;
      } else {
        return `${parentPath}/${route.path}`;
      }
    }
  }
}

// 是否有子路由
function hasChildren(route) {
  if (!route.children) {
    return false;
  } else {
    if (route.alwaysShow) {
      return true;
    }
    return route.children.filter((it) => !it.meta?.hidden).length >= 1;
  }
}

// 点击菜单
function handleMenuItemClick(item) {
  router.push(item.path);
}
</script>

<style scoped lang="scss">
.menu-icon {
  height: 40px;
  line-height: 40px;
}

.title {
  font-size: 16px;
}
</style>
