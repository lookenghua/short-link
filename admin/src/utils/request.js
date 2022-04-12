import { message } from "ant-design-vue";
import { useUserStore } from "@/store/modules/user";
import axios from "axios";

const request = axios.create({
  baseURL: import.meta.env.VUE_APP_API,
  timeout: 5000,
});

request.interceptors.request.use(
  function (config) {
    const userStore = useUserStore();
    const token = userStore.getToken;
    if (token) {
      config.headers.authorization = "Bearer " + token;
    }
    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

// Add a response interceptor
request.interceptors.response.use(
  function (response) {
    const res = response.data;
    if (res.success) {
      return res;
    }

    return Promise.reject(res.errorMessage);
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    message.error(error.toString() || "网络异常");
    return Promise.reject(error);
  }
);

export { request };
