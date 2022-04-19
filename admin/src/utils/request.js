import URI from "urijs";
const defaultOptions = {
  baseURL: String(import.meta.env.VITE_API_URL),
};

/**
 * 合并地址
 * @param {string} url  - 接口地址
 * @param {Object} params  - 参数
 * @return {string}
 */
function mergeURI(url, params = {}) {
  if (url.startsWith("http")) {
    return url;
  }
  const u = new URI(new URL(url, defaultOptions.baseURL).toString());
  for (let key in params) {
    u.addQuery(key, params[key]);
  }
  return u.toString();
}
// 参数合并
function mergeOptions(config = {}, body) {
  return Object.assign({}, config, { body });
}
const request = {
  get(url, config = {}) {
    const serverURI = mergeURI(url, config.params);
    const newConfig = mergeOptions(config);
    return fetch(serverURI, { method: "get", ...newConfig }).then((res) => {
      if (res.status === 200) {
        return res.json();
      } else {
        Promise.reject(res);
      }
    });
  },
  post(url, data, config = {}) {
    const serverURI = mergeURI(url, config.params);
    const newConfig = mergeOptions(config, data);
    return fetch(serverURI, { method: "post", ...newConfig }).then((res) => {
      if (res.status === 200) {
        return res.json();
      } else {
        Promise.reject(res);
      }
    });
  },
  put() {},
  delete() {},
};

export { request };
