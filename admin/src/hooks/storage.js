import { onMounted, ref, watch } from "vue";
import WebStorage from "@exvu/web-storage";

export function useLocalStorage(key, defaultValue = null) {
  const storage = ref(null);
  const webStorage = new WebStorage({
    encrypt: ["key", "value"],
    pre: "_wc_",
    storage: localStorage,
  });

  watch(storage, (val) => {
    console.log("更新数据");
    if (!webStorage.has(key)) {
      webStorage.set(key, val);
    }
  });

  onMounted(() => {
    if (webStorage.has(key)) {
      storage.value = webStorage.get(key);
    } else {
      if (defaultValue != null) {
        storage.value = defaultValue;
      }
    }
  });

  return storage;
}
