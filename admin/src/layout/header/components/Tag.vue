<template>
  <div>
    <div class="tab" @contextmenu="handleContextMenu">
      <span>{{ info.meta?.title }} </span>
      <Icon
        v-if="!info.meta?.affix"
        icon="carbon:close"
        color="white"
        width="16"
        height="16"
        class="icon-close"
        @click.stop="handleTagMenuItemClick('close')"
      />
    </div>
    <n-dropdown
      placement="bottom-start"
      trigger="manual"
      :x="xRef"
      :y="yRef"
      :options="options"
      :show="showDropdown"
      :render-icon="renderIcon"
      @select="handleTagMenuItemClick"
      @clickoutside="showDropdown = false"
    >
    </n-dropdown>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "TagItem",
});
</script>
<script setup>
import { computed, nextTick, toRefs, ref, h, watch } from "vue";
import { useLayoutStore } from "@/store/modules/layout";
import { Icon } from "@iconify/vue/dist/iconify";

const props = defineProps({
  info: {
    type: Object,
    required: true,
  },
});
const emits = defineEmits(["menu"]);
const { info } = toRefs(props);
const showDropdown = ref(false);
const xRef = ref(0);
const yRef = ref(0);
const layoutStore = useLayoutStore();
const activeKey = computed(() => layoutStore.activeTag);
const isActive = computed(() => layoutStore.activeTag === info.value.fullPath);
const pageIndex = computed(() =>
  layoutStore.visitedViews.findIndex(
    (it) => it.fullPath === info.value.fullPath
  )
);
// 列表
const options = ref([
  {
    label: "重新加载",
    key: "reload",
    disabled: false,
  },
  {
    label: "关闭标签页",
    key: "close",
    disabled: false,
  },
  {
    type: "divider",
    key: "divider1",
  },
  {
    label: "关闭左侧标签页",
    key: "closeLeft",
    disabled: false,
  },
  {
    label: "关闭右侧标签页",
    key: "closeRight",
    disabled: false,
  },
  {
    type: "divider",
    key: "divider2",
  },
  {
    label: "关闭其他标签页",
    key: "closeOther",
    disabled: false,
  },
  {
    label: "关闭全部标签页",
    key: "closeAll",
  },
]);
// 是否可以重新加载
const canReloadPage = computed(() => activeKey.value === info.value.fullPath);
// 是否可以关闭页面
const canClosePage = computed(() => !info.value.meta?.affix);
// 是否可以关闭左侧菜单
const canCloseLeftPage = computed(
  () =>
    layoutStore.visitedViews.some(
      (it, i) => i < pageIndex.value && !it.meta?.affix
    ) && isActive.value
);
// 是否可以关闭右侧菜单
const canCloseRightPage = computed(
  () =>
    layoutStore.visitedViews.some(
      (it, i) => i > pageIndex.value && !it.meta?.affix
    ) && isActive.value
);

// 监听多个状态
watch(
  () => [
    canReloadPage,
    canClosePage,
    canCloseLeftPage,
    canCloseRightPage,
    isActive,
  ],
  () => {
    options.value[0].disabled = !canReloadPage.value;
    options.value[1].disabled = !canClosePage.value;
    options.value[3].disabled = !canCloseLeftPage.value;
    options.value[4].disabled = !canCloseRightPage.value;
    options.value[6].disabled = !isActive.value;
  },
  { immediate: true }
);

// 关闭菜单
function handleTagMenuItemClick(directive) {
  emits("menu", directive);
  showDropdown.value = false;
}

// 右键
function handleContextMenu(e) {
  e.preventDefault();
  showDropdown.value = false;
  nextTick().then(() => {
    showDropdown.value = true;
    xRef.value = e.clientX;
    yRef.value = e.clientY;
  });
}

// 渲染图标
function renderIcon(option) {
  let icon = null;
  if (option.key === "reload") {
    icon = {
      icon: "ant-design:reload-outlined",
      color: "rgba(0, 0, 0, 0.85)",
      width: 18,
      height: 18,
      inline: true,
    };
    if (option.disabled) {
      icon.color = "rgba(0, 0, 0, 0.25)";
    }
  }
  if (option.key === "close") {
    icon = {
      icon: "ant-design:close-outlined",
      color: "rgba(0, 0, 0, 0.85)",
      width: 18,
      height: 18,
      inline: true,
    };
    if (option.disabled) {
      icon.color = "rgba(0, 0, 0, 0.25)";
    }
  }
  if (option.key === "closeLeft") {
    icon = {
      icon: "icon-park-outline:to-left",
      color: "rgba(0, 0, 0, 0.85)",
      width: 18,
      height: 18,
      inline: true,
    };
    if (option.disabled) {
      icon.color = "rgba(0, 0, 0, 0.25)";
    }
  }
  if (option.key === "closeRight") {
    icon = {
      icon: "icon-park-outline:to-right",
      color: "rgba(0, 0, 0, 0.85)",
      width: 18,
      height: 18,
      inline: true,
    };
  }
  if (option.key === "closeOther") {
    icon = {
      icon: "icon-park-outline:distribute-vertically",
      color: "rgba(0, 0, 0, 0.85)",
      width: 18,
      height: 18,
      inline: true,
    };
    if (option.disabled) {
      icon.color = "rgba(0, 0, 0, 0.25)";
    }
  }
  if (option.key === "closeAll") {
    icon = {
      icon: "ci:line-l",
      rotate: 1,
      color: "#333",
      width: 18,
      height: 18,
      inline: true,
    };
  }
  return h(Icon, icon);
}
</script>

<style scoped lang="scss">
.active-tab {
  .tab {
    background: #42b983;
    color: white;
    border-color: #42b983;
  }
}

.tab {
  height: 28px;
  padding-right: 10px;
  padding-left: 10px;
  display: flex;
  align-items: center;
  border: 1px solid #d9d9d9;
  margin-right: 4px;
  cursor: pointer;

  .icon-close {
    line-height: 1;
    margin-left: 2px;

    &:hover {
      background-color: #b4bccc;
      color: #fff;
      border-radius: 50%;
    }
  }
}
</style>
