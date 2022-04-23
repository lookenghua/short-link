export function isFormData(v) {
  return Object.prototype.toString.call(v) === "[object FormData]";
}
