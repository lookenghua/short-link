import { request } from "@/utils/request";

// 发布文章
export function createArticle(data) {
  return request.post("/article", data);
}

// 获取文章
export function getArticles(data) {
  return request.get("/article", { params: data });
}

// 文章详情
export function getArticleDetail(id) {
  return request.get(`/article/${id}`);
}

// 删除文章
export function deleteArticle(id) {
  return request.delete(`/article/${id}`);
}

// 编辑文章
export function updateArticle(id, data) {
  return request.put(`/article/${id}`, data);
}

// 获取待审核文章
export function getCommentPendingApproval(params) {
  return request.get("/article/comment/pending-approval", { params });
}

// 审批文章
export function createCommentApprove(commentId, data) {
  return request.post(`/article/comment/${commentId}/approval`, data);
}
