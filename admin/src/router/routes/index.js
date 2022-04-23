const modulesFiles = import.meta.globEager("./modules/*.js");
const asyncRoutes = [];

Object.keys(modulesFiles).forEach((it) => {
  const value = modulesFiles[it];
  const routes = value.default;
  if (Array.isArray(routes)) {
    routes.forEach((route) => {
      asyncRoutes.push(route);
    });
  } else {
    asyncRoutes.push(routes);
  }
});

export { asyncRoutes };

const LoginRoute = {
  path: "/login",
  name: "LoginPage",
  component: () => import("@/pages/system/login/Login.vue"),
  meta: {
    title: "登录页",
  },
};

// eslint-disable-next-line no-unused-vars
const PageNotFoundRoute = {};
export const basicRoutes = [LoginRoute];
