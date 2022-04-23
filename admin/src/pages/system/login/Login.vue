<template>
  <div class="page">
    <img src="../../../assets/images/login/blog.png" alt="" class="img-blog" />
    <div class="login-area">
      <n-card
        title="欢迎登录短链接管理系统"
        :header-style="headStyle"
        :content-style="bodyStyle"
        class="card"
        size="large"
      >
        <n-form
          ref="formRef"
          :model="modelRef"
          :rules="rulesRef"
          size="small"
          :show-label="false"
        >
          <n-form-item path="username">
            <n-input
              v-model:value="modelRef.username"
              placeholder="请输入用户名"
              size="large"
            >
              <template #prefix>
                <Icon
                  icon="ant-design:user-outlined"
                  width="18"
                  height="18"
                  :inline="true"
                  color="rgba(0, 0, 0, 0.25)"
                />
              </template>
            </n-input>
          </n-form-item>
          <n-form-item path="password">
            <n-input
              v-model:value="modelRef.password"
              type="password"
              size="large"
              show-password-on="mousedown"
              placeholder="请输入密码"
            >
              <template #prefix>
                <Icon
                  icon="ant-design:lock-outlined"
                  width="18"
                  height="18"
                  :inline="true"
                  color="rgba(0, 0, 0, 0.25)"
                />
              </template>
            </n-input>
          </n-form-item>
          <n-form-item>
            <n-checkbox v-model:checked="rememberPass">记住密码</n-checkbox>
          </n-form-item>
          <n-form-item>
            <n-button
              type="info"
              block
              round
              size="large"
              :loading="loading"
              @click="submit"
              >登录</n-button
            >
          </n-form-item>
        </n-form>
      </n-card>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from "vue";
import { useLocalStorage } from "@/hooks/storage.js";
import init, { md5encrypt } from "wasm-utils";
import { useUserStore } from "@/store/modules/user";
import { useRouter } from "vue-router";

const formRef = ref(null);
const rememberPass = ref(false);
const loading = ref(false);

init();
const userStore = useUserStore();
const router = useRouter();
const accountStorage = useLocalStorage("account");
const headStyle = {
  border: "none",
  fontSize: "30px",
  fontWeight: "bold",
  paddingTop: "30px",
  textAlign: "center",
};
const bodyStyle = {
  padding: "20px 30px",
};
const modelRef = reactive({
  username: "admin",
  password: "123456",
});
const rulesRef = reactive({
  username: [
    {
      required: true,
      message: "请输入用户名",
    },
  ],
  password: [
    {
      required: true,
      message: "请输入密码",
    },
  ],
});

onMounted(() => {
  if (accountStorage.value !== null) {
    rememberPass.value = true;
    modelRef.username = accountStorage.value.username;
    modelRef.password = accountStorage.value.password;
  }
});

// 校验
function submit(e) {
  e.preventDefault();
  formRef.value?.validate((errors) => {
    if (!errors) {
      const data = {
        username: modelRef.username,
        password: md5encrypt(modelRef.password),
      };
      userStore.login(data).subscribe(
        () => {
          router.push("/");
        },
        (e) => {
          console.log(e);
        }
      );
    }
  });
}
</script>

<style scoped lang="scss">
.page {
  width: 100vw;
  height: 100vh;
  position: relative;
  background: url("../../../assets/images/login/bg.png") center center no-repeat;
  background-size: 100% 100%;
}

.login-area {
  position: absolute;
  top: 50%;
  right: 198px;
  transform: translateY(-60%);
  width: 470px;
}

.img-blog {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 600px;
  left: 100px;
}

.card {
  border-radius: 10px;
}

.input {
  border-radius: 5px;
}

@media (max-width: 1500px) {
  .img-blog {
    display: none;
  }

  .login-area {
    right: 50%;
    transform: translate(50%, -60%);
  }
}
</style>
