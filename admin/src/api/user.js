import { request } from "@/utils/request";

// 管理员登录
export function userLogin(data) {
  return request.post("/v1/user/token", data);
}

// 用户信息
export function getCurrentAdminInfo() {
  return request.get(`/v1/user`);
}
