import { TOKEN_KEY } from "../enums/cache.enum";

// 获取token
export function getToken() {
  return localStorage.getItem(TOKEN_KEY);
}
