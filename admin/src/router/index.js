import { basicRoutes } from "./routes";
import { createRouter, createWebHistory } from "vue-router";

const WHITE_NAME_LIST = [];

// 获取路由名称
function getRouteNames(array) {
  array.forEach((item) => {
    WHITE_NAME_LIST.push(item.name);
    getRouteNames(item.children || []);
  });
}

getRouteNames(basicRoutes);

// 创建路由
export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: basicRoutes,
  strict: true,
  scrollBehavior: (to, from, savedPosition) => {
    if (savedPosition) {
      return savedPosition;
    } else {
      return {
        top: 0,
      };
    }
  },
});

// 重置路由
export function resetRouter() {
  router.getRoutes().forEach((route) => {
    const { name } = route;
    if (name && !WHITE_NAME_LIST.includes(name)) {
      router.hasRoute(name) && router.removeRoute(name);
    }
  });
}

// 设置路由
export function setupRouter(app) {
  app.use(router);
}
