import { request } from "@/utils/request";

// 管理员登录
export function login(data) {
  return request.post("/admin/token", data);
}

// 用户信息
export function getCurrentAdminInfo() {
  return request.get("/admin/info");
}
