import message from "ant-design-vue/lib/message";

export function useLoading() {
  const key = new Date().getTime();

  message.config({ top: "20vh" });

  // 显示loading
  function showLoading(content) {
    message.loading({ content, key });
  }

  // 关闭loading
  function closeLoading() {
    message.destroy();
  }

  // 加载成功
  function loadingSuccess(content, duration = 2) {
    message.success({ content, key, duration });
  }

  // 加载失败
  function loadingError(content, duration = 2) {
    message.error({ content, key, duration });
  }

  return {
    showLoading,
    closeLoading,
    loadingSuccess,
    loadingError,
  };
}
