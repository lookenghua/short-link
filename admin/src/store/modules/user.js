import { defineStore } from "pinia";
import { userLogin } from "@/api/user";
import { useLocalStorage, useSessionStorage } from "@vueuse/core";
import { tap } from "rxjs";

export const useUserStore = defineStore("user", {
  state: () => ({
    userInfo: useSessionStorage("userInfo", {}), // 用户信息
    token: useLocalStorage("token", ""), // token
  }),
  getters: {
    // 获取token
    getToken(state) {
      return state.token;
    },
    avatar(state) {
      return state.userInfo?.avatar;
    },
    username(state) {
      return state.userInfo?.username;
    },
  },
  actions: {
    // 登录
    login(data) {
      return userLogin(data).pipe(
        tap((res) => {
          this.token = res.data;
        })
      );
    },
    // 设置管理员信息
    setUserInfo(data) {
      console.log(data);
      this.userInfo = data;
    },
    // 退出登录
    logout() {
      this.userInfo = null;
      this.token = "";
    },
  },
});
