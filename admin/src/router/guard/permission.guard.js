import "nprogress/nprogress.css";
import { asyncRoutes } from "../routes";
import { getCurrentAdminInfo } from "@/api/user";
import { isEmpty } from "@/utils/util.js";
import { isNavigationFailure } from "vue-router";
import { Logger } from "@/utils/logger.js";
import { usePermissionStore } from "@/store/modules/perssion";
import { useUserStore } from "@/store/modules/user";
import NProgress from "nprogress";

// 路由权限守卫
export function createPermissionGuard(router) {
  const userStore = useUserStore();
  const permissionStore = usePermissionStore();
  router.beforeEach(async (to, from, next) => {
    const token = userStore.getToken;
    Logger.i("是否有token:" + !!token);
    NProgress.start();
    if (!token) {
      if (to.path === "/login") {
        next(true);
      } else {
        next({ path: "/login" });
      }
    } else {
      NProgress.inc();
      const userInfo = userStore.userInfo;
      if (isEmpty(userInfo)) {
        console.log("===正在获取用户信息");
        try {
          NProgress.inc();
          const res = await getCurrentAdminInfo().toPromise();
          userStore.setUserInfo(res.data);
        } catch (e) {
          console.error(e);
        } finally {
          NProgress.inc();
        }
      }

      if (!permissionStore.isDynamicAddRoute) {
        NProgress.inc();
        asyncRoutes.forEach((route) => {
          router.addRoute(route);
        });
        permissionStore.setDynamicAddRoute(true);
        permissionStore.generatePermissionRoutes(asyncRoutes);
        console.log("====添加路由成功====");
        next({ ...to, replace: true });
        NProgress.inc();
      } else {
        if (to.path === "/login" || to.path === "/") {
          next({ path: "/dashboard" });
        } else {
          next(true);
        }
        NProgress.inc();
      }
    }
  });
  router.afterEach((to, from, failure) => {
    if (isNavigationFailure(failure)) {
      console.log("failed navigation", failure);
    }
    NProgress.done();
  });
}
