import { reactive } from 'vue'

/**
 * reactive封装
 * @template T
 * @param { T } defaultData
 * @returns { {data:T, reset:function, set:function} }
 */
export function useReactive(defaultData = {}) {
  const data = reactive({ ...defaultData })

  // 重置
  function reset() {
    Object.assign(data, defaultData)
  }

  //设置
  function set(newData = {}) {
    Object.assign(data, newData)
  }

  return { data, reset, set }
}
