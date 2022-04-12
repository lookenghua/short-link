import { computed, ref } from 'vue'

/**
 * 加密字符串
 * @param {string} code
 * @return {string}
 */
function compileStr(code) {
  let c = String.fromCharCode(code.charCodeAt(0) + code.length)
  for (let i = 1; i < code.length; i++) {
    c += String.fromCharCode(code.charCodeAt(i) + code.charCodeAt(i - 1))
  }
  return escape(c) + '_c'
}

/**
 * 解密字符串
 * @param {string} code
 * @return {string}
 */
function uncompileStr(code) {
  code = unescape(code.substring(0, code.length - 2))
  let c = String.fromCharCode(code.charCodeAt(0) - code.length)
  for (let i = 1; i < code.length; i++) {
    c += String.fromCharCode(code.charCodeAt(i) - c.charCodeAt(i - 1))
  }
  return c
}
/**
 * @typedef {Object} StorageOptions
 * @property {Storage} storage - 存储位置
 */
const defaultOptions = {
  storage: sessionStorage,
}

/**
 * 安全存到储sessionStorage
 * @param {string} key - 存储的key
 * @param {T} initialValue - 初始值
 * @param {StorageOptions} options - 配置
 * @returns {WritableComputedRef<T>}
 */
function useStorage(key, initialValue, options = defaultOptions) {
  let data = ref(initialValue)
  try {
    const info = options.storage.getItem(key)
    if (info) {
      data.value = JSON.parse(uncompileStr(info))
    }
  } catch (e) {
    console.log(e)
  }

  return computed({
    get() {
      return data
    },
    set(value) {
      data.value = value
      if (value === null) {
        options.storage.removeItem(key)
      } else {
        const data = compileStr(JSON.stringify(value))
        options.storage.setItem(key, data)
      }
    },
  })
}

/**
 * 安全存到储sessionStorage
 * @param {string} key - 存储的key
 * @param {T} initialValue - 初始值
 * @return {WritableComputedRef<T>}
 */
export function useSecureSessionsStorage(key, initialValue = '') {
  return useStorage(key, initialValue, { storage: sessionStorage })
}
