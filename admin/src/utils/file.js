// 拼接文件链接
export function completeLink(link) {
  if (!link) return "";
  if (link.startsWith("http") || link.startsWith("//")) {
    return link;
  } else {
    if (link.startsWith("/")) {
      return import.meta.env.VUE_APP_FILE_DOMAIN + link;
    } else {
      return import.meta.env.VUE_APP_FILE_DOMAIN + "/" + link;
    }
  }
}
