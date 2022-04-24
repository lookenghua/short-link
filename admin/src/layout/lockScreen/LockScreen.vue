<template>
  <div
    class="fixed top-0 left-0 w-full h-full z-50 flex items-center justify-center"
    style="background-color: #151515"
  >
    <div
      class="absolute left-1/2 top-3 text-center cursor-pointer"
      @click="showModal = true"
    >
      <Icon icon="akar-icons:lock-on" color="white" width="30" height="30" />
      <div class="text-white">点击解锁</div>
    </div>
    <FlipClock />
    <div
      v-if="showModal"
      class="absolute top-0 left-0 w-full h-full bg-black bg-opacity-80 z-30 flex items-center justify-center"
    >
      <div class="relative top--4 flex flex-col items-center justify-center">
        <OpenData size="20" type="avatar" class="align-text-bottom" />
        <OpenData type="username" class="text-lg text-white mb-3" />
        <n-form
          label-placement="left"
          :model="dialogFormRef"
          :rules="rulesRef"
          ref="formRef"
        >
          <n-form-item path="password">
            <n-input
              type="password"
              v-model:value="dialogFormRef.password"
              placeholder="请输入锁屏密码"
              show-password-on="mousedown"
              @focus="clearValidate"
            />
          </n-form-item>
        </n-form>
        <div class="flex items-center justify-between w-full mt-2">
          <span class="text-blue-600 cursor-pointer" @click="showModal = false"
            >返回</span
          >
          <span class="text-blue-600 cursor-pointer" @click="logout"
            >返回登录</span
          >
          <span class="text-blue-600 cursor-pointer" @click="unlock"
            >进入系统</span
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "LockScreen",
});
</script>
<script setup>
import { onMounted, reactive, ref } from "vue";
import { useAppStore } from "@/store/modules/app.js";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/modules/user.js";
import FlipClock from "@/components/FlipClock";

const formRef = ref(null);
const appStore = useAppStore();
const router = useRouter();
let showModal = ref(false);
const dialogFormRef = reactive({
  password: "",
});
const rulesRef = reactive({
  password: [
    {
      validator: async (rule, value) => {
        if (!value) {
          return Promise.reject("请输入锁屏密码");
        } else if (value !== appStore.lockScreenPassword) {
          return Promise.reject("锁屏密码不正确");
        }
        return Promise.resolve();
      },
    },
  ],
});
// 解锁
function unlock() {
  formRef.value?.validate().then((errors) => {
    if (!errors) {
      appStore.lockScreenPassword = null;
      clearValidate();
    }
  });
}

// 清除校验信息
function clearValidate() {
  formRef.value?.restoreValidation();
}

// 返回登录
function logout() {
  const userStore = useUserStore();
  userStore.logout();
  appStore.lockScreenPassword = null;
  router.push("/login");
}

onMounted(() => {});
</script>
