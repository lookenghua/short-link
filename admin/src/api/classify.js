import { request } from "@/utils/request";

// 获取分类
export function getClassifies(data) {
  return request.get("/classify", { params: data });
}

// 创建分类
export function createClassify({ name = "", parentId = null }) {
  return request.post("/classify", { name, parentId });
}
