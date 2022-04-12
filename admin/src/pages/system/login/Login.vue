<template>
  <div class="page">
    <img src="../../../assets/images/login/blog.png" alt="" class="img-blog" />
    <div class="login-area">
      <a-card
        title="欢迎登录博客管理系统"
        :head-style="headStyle"
        :body-style="bodyStyle"
        class="card"
      >
        <a-form>
          <a-form-item>
            <a-input
              v-model:value="modelRef.username"
              size="large"
              placeholder="请输入用户名"
              class="input"
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
            </a-input>
          </a-form-item>
          <a-form-item style="margin-bottom: 10px">
            <a-input-password
              v-model:value="modelRef.password"
              type="password"
              size="large"
              placeholder="请输入密码"
              class="input"
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
            </a-input-password>
          </a-form-item>
          <a-form-item>
            <a-checkbox v-model:checked="rememberPass">记住密码</a-checkbox>
          </a-form-item>
          <a-form-item>
            <a-button
              type="primary"
              block
              shape="round"
              size="large"
              :loading="loading"
              @click="submit"
            >
              登录
            </a-button>
          </a-form-item>
        </a-form>
      </a-card>
    </div>
  </div>
</template>

<script setup>
import { encryptData } from "@/api/tool.js";
import { Form, message } from "ant-design-vue";
import { onMounted, reactive, ref, toRaw } from "vue";
import { useLocalStorage } from "@/hooks/storage.js";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/modules/user.js";

const useForm = Form.useForm;
const rememberPass = ref(false);
const loading = ref(false);

const accountStorage = useLocalStorage("account");
const headStyle = {
  border: "none",
  fontSize: "30px",
  fontWeight: "bold",
  paddingTop: "30px",
  textAlign: "center",
};
const bodyStyle = {
  padding: "35px",
};
const modelRef = reactive({
  username: "",
  password: "",
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
const { validate } = useForm(modelRef, rulesRef);
const router = useRouter();

onMounted(() => {
  if (accountStorage.value !== null) {
    rememberPass.value = true;
    modelRef.username = accountStorage.value.username;
    modelRef.password = accountStorage.value.password;
  }
});

// 校验
function submit() {
  validate().then(() => {
    loading.value = true;
    encryptData(modelRef.password)
      .then((res) => {
        const userStore = useUserStore();
        userStore
          .login({ username: modelRef.username, password: res.data })
          .then(() => {
            if (rememberPass.value) {
              accountStorage.value = toRaw(modelRef);
            }
            loading.value = false;
            message.success("登录成功");
            router.push("/dashboard");
          });
      })
      .catch(() => {
        loading.value = false;
      });
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
