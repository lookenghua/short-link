import { defineStore } from "pinia";

export const usePermissionStore = defineStore("permission", {
  state: () => ({
    isDynamicAddRoute: false, // 是否动态添加路由
    permissionRoutes: [], // 权限路由列表
  }),
  getters: {
    // 获取权限路由列表
    getPermissionRoutes(state) {
      return state.permissionRoutes;
    },
  },
  actions: {
    // 设置是否添加路由状态
    setDynamicAddRoute(added) {
      this.isDynamicAddRoute = added;
    },
    // 生成权限路由
    generatePermissionRoutes(routes) {
      this.permissionRoutes = routes;
    },
  },
});
