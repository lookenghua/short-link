import { request } from "../utils/request";

// 加密数据
export function encryptData(str) {
  return request.post("/tool/encrypt", { data: str });
}

// 上传文件
export function uploadFile(file, annex, folder) {
  const formData = new FormData();
  formData.append("file", file);
  if (annex) {
    formData.append("annex", annex);
  }
  if (folder) {
    formData.append("folder", folder);
  }
  return request.post("/tool/file", formData);
}

// 上传Base64字符串
export function uploadBase64File(file, filename, annex, folder) {
  return request.post("/tool/base64", { file, filename, annex, folder });
}

// 上传自定义文件
export function uploadCustomFile(file, annex, folder) {
  return new Promise((resolve, reject) => {
    if (file.fileType === "http") {
      resolve(file.origin);
    } else if (file.fileType === "file") {
      uploadFile(file.origin, annex, folder)
        .then((data) => {
          resolve(data);
        })
        .catch((e) => {
          reject(e);
        });
    } else if (file.fileType === "base64") {
      uploadBase64File(file.origin, file.name, annex, folder)
        .then((data) => {
          resolve(data);
        })
        .catch((e) => {
          reject(e);
        });
    }
  });
}
