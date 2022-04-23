import "./assets/css/global.css";
import "normalize.css/normalize.css";
import "uno.css";
import { createApp } from "vue";
import { router, setupRouter } from "./router";
import { setupGlobComponents } from "./components";
import { setupRouterGuard } from "./router/guard";
import { setupStore } from "./store";
import App from "./App.vue";

async function bootstrap() {
  const app = createApp(App);

  // 状态管理
  setupStore(app);

  // 全局组件
  setupGlobComponents(app);

  // 路由
  setupRouter(app);

  // 路由守卫
  setupRouterGuard(router);

  // 等待路由完成
  await router.isReady();

  // 挂载DOM
  app.mount("#app");
}

bootstrap().then();
