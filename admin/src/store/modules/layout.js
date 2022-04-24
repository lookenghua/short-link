import { defineStore } from "pinia";
import { Logger } from "@/utils/logger.js";
import { useSessionStorage } from "@vueuse/core";

export const useLayoutStore = defineStore("layout", {
  state: () => ({
    collapsed: false, // 是否缩小
    visitedViews: useSessionStorage("visitedViews", []), // 访问过的页面
    cachedViews: [], // 缓存的页面
    activeTag: "", // 选中的标签
    activeMenu: [""], // 选中的菜单
    openKeys: "", //当前展开的 SubMenu 菜单项 key 数组
  }),
  getters: {
    lastVisitedView(state) {
      return state.visitedViews.length === 0
        ? null
        : state.visitedViews[state.visitedViews.length - 1];
    },
    activeVisitedIndex(state) {
      return state.visitedViews.findIndex(
        (item) => item.fullPath === state.activeTag
      );
    },
    activeView(state) {
      return state.visitedViews.find(
        (item) => item.fullPath === state.activeTag
      );
    },
  },
  actions: {
    // 添加视图
    addView(view) {
      this.activeMenu = [view.fullPath];
      this.activeTag = view.fullPath;
      if (this.visitedViews.some((it) => it.fullPath === view.fullPath)) {
        Logger.d("===已存在===");
        return;
      }
      Logger.d(view);
      if (view.meta.affix) {
        const i = this.visitedViews.findIndex((it) => !it.meta.affix);
        if (i < 0) {
          this.visitedViews.push(view);
        } else {
          this.visitedViews.splice(i, 0, view);
        }
      } else {
        this.visitedViews.push(view);
      }
    },
    // 移除视图
    removeView(view) {
      const fullPathList = this.visitedViews.map((it) => it.fullPath);
      const i = fullPathList.indexOf(view.fullPath);
      console.log(i);
      if (i !== -1) {
        this.visitedViews.splice(i, 1);
        console.log(this.visitedViews);
      }
    },
    // 设置选中的标签栏
    setActiveTag(tag) {
      this.activeTag = tag;
    },
    // 设置打开的菜单
    setActiveMenu(menu) {
      this.activeMenu = menu;
    },
    // 设置
    setVisitedViews(views) {
      this.visitedViews = views;
    },
  },
});
