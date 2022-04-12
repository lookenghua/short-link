import { Modal } from "ant-design-vue";

// 显示确认框
export function showConfirm({
  title = "",
  icon = null,
  content = null,
  okText = "确定",
  okType = "primary",
}) {
  return new Promise((resolve, reject) => {
    Modal.confirm({
      title,
      icon,
      content,
      okText,
      okType,
      onOk() {
        resolve();
      },
      onCancel() {
        reject();
      },
    });
  });
}
// 格式化代码
export function formatBytes(bytes, decimals) {
  if (bytes === 0) return "0 Bytes";
  const k = 1024,
    dm = decimals || 2,
    sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"],
    i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
}
// 是否为空
export function isEmpty(data) {
  if (!data) {
    return true;
  }
  if (Array.isArray(data)) {
    return data.length === 0;
  }
  if (typeof data === "object") {
    return Object.keys(data).length === 0;
  }
  return false;
}
