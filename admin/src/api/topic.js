import { request } from "../utils/request";

// 分页查询专题
export function getTopics(data) {
  return request.get("/topic", { params: data });
}

// 创建分类
export function createTopic({ title = "", description = "", cover = "" }) {
  return request.post("/topic", { title, description, cover });
}

// 删除分类
export function deleteTopic(id) {
  return request.delete(`/topic/${id}`);
}

// 更新分类
export function updateTopic(id, data) {
  return request.put(`/topic/${id}`, data);
}
