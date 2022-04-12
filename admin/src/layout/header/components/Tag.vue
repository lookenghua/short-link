<template>
  <a-dropdown :trigger="['contextmenu']" overlay-class-name="right-menu">
    <div class="tab">
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
    <template #overlay>
      <a-menu>
        <a-menu-item
          key="reload"
          :disabled="!canReloadPage"
          @click="handleTagMenuItemClick('reload')"
        >
          <div class="inline-flex items-center">
            <Icon
              v-if="!canReloadPage"
              icon="ant-design:reload-outlined"
              color="rgba(0, 0, 0, 0.25)"
              width="18"
              height="18"
              :inline="true"
            />
            <Icon
              v-else
              icon="ant-design:reload-outlined"
              color="rgba(0, 0, 0, 0.85)"
              width="18"
              height="18"
              :inline="true"
            />

            <span class="ml-1">重新加载</span>
          </div>
        </a-menu-item>
        <a-menu-item
          key="close"
          :disabled="!canClosePage"
          @click="handleTagMenuItemClick('close')"
        >
          <div class="inline-flex items-center">
            <Icon
              v-if="!canClosePage"
              icon="ant-design:close-outlined"
              color="rgba(0, 0, 0, 0.25)"
              width="18"
              height="18"
              :inline="true"
            />
            <Icon
              v-else
              icon="ant-design:close-outlined"
              color="rgba(0, 0, 0, 0.85)"
              width="18"
              height="18"
              :inline="true"
            />
            <span class="ml-1">关闭标签页</span>
          </div>
        </a-menu-item>
        <a-menu-divider />
        <a-menu-item
          key="closeLeft"
          :disabled="!canCloseLeftPage"
          @click="handleTagMenuItemClick('closeLeft')"
        >
          <div class="inline-flex items-center">
            <Icon
              v-if="!canCloseLeftPage"
              icon="icon-park-outline:to-left"
              color="rgba(0, 0, 0, 0.25)"
              width="18"
              height="18"
              :inline="true"
            />
            <Icon
              v-else
              icon="icon-park-outline:to-left"
              color="rgba(0, 0, 0, 0.85)"
              width="18"
              height="18"
              :inline="true"
            />
            <span class="ml-1">关闭左侧标签页</span>
          </div>
        </a-menu-item>
        <a-menu-item
          key="closeRight"
          :disabled="!canCloseRightPage"
          @click="handleTagMenuItemClick('closeRight')"
        >
          <div class="inline-flex items-center">
            <Icon
              v-if="!canCloseRightPage"
              icon="icon-park-outline:to-right"
              color="rgba(0, 0, 0, 0.25)"
              width="18"
              height="18"
              :inline="true"
            />
            <Icon
              v-else
              icon="icon-park-outline:to-right"
              color="rgba(0, 0, 0, 0.85)"
              width="18"
              height="18"
              :inline="true"
            />
            <span class="ml-1">关闭右侧标签页</span>
          </div>
        </a-menu-item>
        <a-menu-divider />
        <a-menu-item
          key="closeOther"
          :disabled="!isActive"
          @click="handleTagMenuItemClick('closeOther')"
        >
          <div class="inline-flex items-center">
            <Icon
              v-if="!isActive"
              icon="icon-park-outline:distribute-vertically"
              color="rgba(0, 0, 0, 0.25)"
              width="18"
              height="18"
              :inline="true"
            />
            <Icon
              v-else
              icon="icon-park-outline:distribute-vertically"
              color="rgba(0, 0, 0, 0.85)"
              width="18"
              height="18"
              :inline="true"
            />
            <span class="ml-1"> 关闭其他标签页</span>
          </div>
        </a-menu-item>
        <a-menu-item key="closeAll" @click="handleTagMenuItemClick('closeAll')">
          <div class="inline-flex items-center">
            <Icon
              icon="ci:line-l"
              :rotate="1"
              color="#333"
              width="18"
              height="18"
              :inline="true"
            />
            <span class="ml-1">关闭全部标签页</span>
          </div>
        </a-menu-item>
      </a-menu>
    </template>
  </a-dropdown>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "TagPage",
});
</script>
<script setup>
import { computed, toRefs } from "vue";
import { useLayoutStore } from "@/store/modules/layout";

const props = defineProps({
  info: {
    type: Object,
    required: true,
  },
});
const emits = defineEmits(["menu"]);
const { info } = toRefs(props);
const layoutStore = useLayoutStore();
const activeKey = computed(() => layoutStore.activeTag);
const isActive = computed(() => layoutStore.activeTag === info.value.fullPath);
const pageIndex = computed(() =>
  layoutStore.visitedViews.findIndex(
    (it) => it.fullPath === info.value.fullPath
  )
);
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
// 关闭菜单
function handleTagMenuItemClick(directive) {
  emits("menu", directive);
}
</script>

<style scoped lang="scss">
.tab {
  height: 28px;
  padding-right: 10px;
  padding-left: 10px;
  display: flex;
  align-items: center;
  border: 1px solid #d9d9d9;
  margin-right: 4px;
  cursor: pointer;

  &.active-tab {
    background: #42b983;
    color: white;
    border-color: #42b983;
  }

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
