const TOKEN_KEY = "token";

// 获取token
export function getToken() {
  return localStorage.getItem(TOKEN_KEY);
}
