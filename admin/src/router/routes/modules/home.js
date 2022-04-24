const Layout = () => import("@/layout/index.vue");
const homeRoute = {
  path: "/",
  component: Layout,
  name: "Home",
  meta: {
    title: "首页",
    icon: "mdi-light:home",
    sort: 100,
  },
  children: [
    {
      path: "/dashboard",
      name: "Dashboard",
      meta: {
        title: "Dashboard",
        icon: "ant-design:dashboard-outlined",
        sort: 101,
        affix: true,
      },
      component: () => import("@/pages/home/dashboard/dashboard.vue"),
    },
  ],
};

export default homeRoute;
