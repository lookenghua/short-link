import { defineStore } from "pinia";
import { useSecureSessionsStorage } from "@/plugin/use-api/storage.js";

export const useAppStore = defineStore("app", {
  state: () => ({
    showRouterView: true,
    lockScreenPassword: useSecureSessionsStorage("lockScreenPassword", ""),
  }),
  getters: {
    lockScreen(state) {
      return !!state.lockScreenPassword;
    },
  },
  actions: {
    reloadPage() {
      this.showRouterView = false;
      setTimeout(() => {
        this.showRouterView = true;
      }, 1000);
    },
  },
});
