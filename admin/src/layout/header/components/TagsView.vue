<template>
  <div class="tags-area">
    <div class="tabs flex items-center">
      <Tag
        v-for="(item, index) in affixVisitedViews"
        :key="item.fullPath"
        :class="{
          'active-tab': item.fullPath === activeKey,
          'affix-tab': item.meta?.affix,
        }"
        :info="item"
        @click="handleMenuItemClick(item.fullPath)"
        @menu="(menu) => handleTagMenuItemClick(index, menu)"
      />
    </div>
    <div ref="tabsRef" class="tabs flex items-center flex-1">
      <Tag
        v-for="(item, index) in unaffixVisitedViews"
        :key="item.fullPath"
        :class="{
          'active-tab': item.fullPath === activeKey,
          'affix-tab': item.meta?.affix,
        }"
        :info="item"
        @click="handleMenuItemClick(item.fullPath)"
        @menu="
          (menu) =>
            handleTagMenuItemClick(affixVisitedViews.length + index, menu)
        "
      />
    </div>
    <div
      class="btn-item"
      @click="handleTagMenuItemClick(activeIndex, 'reload')"
    >
      <Icon
        icon="grommet-icons:power-reset"
        color="#333"
        width="18"
        height="18"
        :inline="true"
      />
    </div>
    <a-dropdown overlay-class-name="right-menu">
      <div class="btn-item">
        <Icon
          icon="ant-design:down-outlined"
          color="#333"
          width="18"
          height="18"
          :inline="true"
        />
      </div>
      <template #overlay>
        <a-menu>
          <a-menu-item
            key="reload"
            @click="handleTagMenuItemClick(activeIndex, 'reload')"
          >
            <div class="inline-flex items-center">
              <Icon
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
            @click="handleTagMenuItemClick(activeIndex, 'close')"
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
            @click="handleTagMenuItemClick(activeIndex, 'closeLeft')"
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
            @click="handleTagMenuItemClick(activeIndex, 'closeRight')"
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
            @click="handleTagMenuItemClick(activeIndex, 'closeOther')"
          >
            <div class="inline-flex items-center">
              <Icon
                icon="icon-park-outline:distribute-vertically"
                color="rgba(0, 0, 0, 0.85)"
                width="18"
                height="18"
                :inline="true"
              />
              <span class="ml-1"> 关闭其他标签页</span>
            </div>
          </a-menu-item>
          <a-menu-item
            key="closeAll"
            @click="handleTagMenuItemClick(activeIndex, 'closeAll')"
          >
            <div class="inline-flex items-center">
              <Icon
                icon="ci:line-l"
                :rotate="1"
                color="rgba(0, 0, 0, 0.85)"
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
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "TagsView",
});
</script>
<script setup>
import { computed, onMounted, ref, toRaw, watch } from "vue";
import { useAppStore } from "@/store/modules/app.js";
import { useLayoutStore } from "@/store/modules/layout";
import { useRoute, useRouter } from "vue-router";
import Sortable from "sortablejs";
import Tag from "./Tag.vue";

const route = useRoute();
const router = useRouter();
const layoutStore = useLayoutStore();
const appStore = useAppStore();

let tabsRef = ref(null);

const visitedViews = computed(() => layoutStore.visitedViews);
const affixVisitedViews = computed(() =>
  layoutStore.visitedViews.filter((it) => it.meta?.affix)
);
const unaffixVisitedViews = computed(() =>
  layoutStore.visitedViews.filter((it) => !it.meta?.affix)
);
const activeIndex = computed(() => layoutStore.activeVisitedIndex);
const activeView = computed(() => layoutStore.activeView);
const canClosePage = computed(() => !activeView.value.meta?.affix);
// 是否可以关闭左侧菜单
const canCloseLeftPage = computed(() =>
  layoutStore.visitedViews.some(
    (it, i) => i < activeIndex.value && !it.meta?.affix
  )
);
// 是否可以关闭左侧菜单
const canCloseRightPage = computed(() =>
  layoutStore.visitedViews.some(
    (it, i) => i > activeIndex.value && !it.meta?.affix
  )
);
let activeKey = computed({
  get() {
    return layoutStore.activeTag;
  },
  set(val) {
    router.push(val);
    layoutStore.setActiveTag(val);
  },
});

watch(
  route,
  (current) => {
    const matched = current.matched;
    const arr = matched
      .filter((it) => it.children.length > 0)
      .map((it) => it.path);
    layoutStore.openKeys = arr;
    if (["/login"].includes(current.path)) {
      return;
    }
    const view = {
      fullPath: current.fullPath,
      meta: current.meta,
      name: current.name,
      params: current.params,
    };
    layoutStore.addView(view);
  },
  { immediate: true }
);

// 监听标签右键菜单
function handleTagMenuItemClick(index, type) {
  if (type === "close") {
    closeMenu(index);
  } else if (type === "reload") {
    appStore.reloadPage();
  } else if (type === "closeLeft") {
    const views = layoutStore.visitedViews;
    layoutStore.visitedViews = views.filter(
      (it, i) => !(i < activeIndex.value && !it.meta?.affix)
    );
  } else if (type === "closeRight") {
    const views = layoutStore.visitedViews;
    layoutStore.visitedViews = views.filter(
      (it, i) => !(activeIndex.value < i && !it.meta?.affix)
    );
  } else if (type === "closeOther") {
    const views = layoutStore.visitedViews;
    layoutStore.visitedViews = views.filter(
      (it, i) => activeIndex.value === i || it.meta?.affix
    );
  } else if (type === "closeAll") {
    layoutStore.visitedViews = [];
    router.push("/");
  }
}

// 关闭菜单
function closeMenu(index) {
  const visitedViews = layoutStore.visitedViews;
  const closeView = visitedViews[index];
  layoutStore.removeView(closeView);
  if (closeView.fullPath === route.fullPath) {
    const lastView = layoutStore.lastVisitedView;
    if (lastView) {
      router.push(lastView.fullPath);
    } else {
      router.push("/");
    }
  }
}

// 点击按钮
function handleMenuItemClick(path) {
  console.log(path);
  activeKey.value = path;
}

onMounted(() => {
  Sortable.create(tabsRef.value, {
    group: "tags",
    handle: ".tab",
    filter: ".affix-tab",
    animation: 150,
    easing: "cubic-bezier(1, 0, 0, 1)",
    preventOnFilter: true,
    removeCloneOnHide: true,
    onEnd(evt) {
      const oldIndex = affixVisitedViews.value.length + evt.oldIndex;
      const newIndex = affixVisitedViews.value.length + evt.newIndex;
      const arr = JSON.parse(JSON.stringify(toRaw(visitedViews.value)));
      if (oldIndex === newIndex) {
        return;
      }
      let oldItem = arr.splice(oldIndex, 1)[0];
      if (oldIndex > newIndex) {
        arr.splice(newIndex, 0, oldItem);
      } else {
        arr.splice(newIndex, 0, oldItem);
      }
      console.log(oldIndex, newIndex, arr);
      layoutStore.visitedViews = arr;
    },
  });
});
</script>

<style scoped lang="scss">
.tags-area {
  background: white;
  padding: 0 0 0 10px;
  border-bottom: 1px solid #d9d9d9;
  display: flex;
  justify-content: space-between;

  .tabs {
    margin: 2px 0;
  }

  .btn-item {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 38px;
    height: 31px;
    color: #00000073;
    text-align: center;
    cursor: pointer;
    border-left: 1px solid #d9d9d9;
  }
}
</style>
