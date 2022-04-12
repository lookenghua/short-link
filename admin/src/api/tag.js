import { request } from "../utils/request";

// 获取标签
export function getTags(data) {
  return request.get("/tag", { params: data });
}

// 创建标签
export function createTag(data) {
  return request.post("/tag", data);
}
