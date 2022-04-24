<template>
  <div class="header flex">
    <div class="flex items-center">
      <div class="icon-item flex items-center justify-center">
        <Icon
          v-if="collapsed"
          icon="ant-design:menu-unfold-outlined"
          width="18"
          height="18"
          class="trigger icon"
          :inline="true"
          @click="() => (collapsed = !collapsed)"
        />
        <Icon
          v-else
          icon="ant-design:menu-fold-outlined"
          width="18"
          height="18"
          :inline="true"
          class="trigger icon"
          @click="() => (collapsed = !collapsed)"
        />
      </div>
      <Breadcrumb />
    </div>
    <div class="flex items-center">
      <div class="icon-item">
        <MessageNotice />
      </div>
      <div class="icon-item">
        <Icon
          v-if="!isFullscreen"
          icon="ant-design:fullscreen-outlined"
          width="24"
          height="24"
          :inline="true"
          class="icon"
          @click="enter"
        />

        <Icon
          v-else
          icon="ant-design:fullscreen-exit-outlined"
          width="24"
          height="24"
          :inline="true"
          class="icon"
          @click="exit"
        />
      </div>
      <div class="icon-item">
        <n-dropdown
          :options="userDropOptions"
          placement="bottom-start"
          trigger="click"
          @select="handleUserMenuClick"
        >
          <div class="user-area flex items-center justify-center h-full">
            <OpenData type="avatar" class="mr-1" />
            <OpenData type="username" class="text-base" />
          </div>
        </n-dropdown>
      </div>
    </div>
  </div>
  <TagsView />
  <n-modal
    v-model:show="lockScreenModalShow"
    preset="card"
    :bordered="false"
    size="huge"
    title="锁定屏幕"
    :style="{ width: '600px' }"
  >
    <div class="pb-8">
      <div class="flex flex-col items-center">
        <OpenData type="avatar" :size="64" class="align-text-bottom" />
        <OpenData type="username" class="text-lg" />
      </div>
      <n-form
        ref="formRef"
        :model="dialogFormRef"
        :rules="rulesRef"
        :show-require-mark="true"
      >
        <n-form-item
          path="password"
          label="锁屏密码"
          require-mark-placement="left"
        >
          <n-input
            v-model:value="dialogFormRef.password"
            placeholder="请输入锁屏密码"
            @focus="clearValidate()"
          />
        </n-form-item>
      </n-form>
      <n-button type="info" block class="mt-3 !rounded-sm" @click="lockScreen">
        锁定
      </n-button>
    </div>
  </n-modal>
  <LockScreen v-if="isLockScreen" />
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({ name: "LayoutHeader" });
</script>
<script setup>
import { computed, createVNode, h, reactive, ref, watch } from "vue";
import { Icon } from "@iconify/vue";
import { Modal } from "ant-design-vue";
import { useAppStore } from "@/store/modules/app.js";
import { useFullscreen } from "@vueuse/core";
import { useLayoutStore } from "@/store/modules/layout";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/modules/user";
import Breadcrumb from "@/layout/header/components/Breadcrumb.vue";
import LockScreen from "@/layout/lockScreen/LockScreen.vue";
import MessageNotice from "@/layout/header/components/MessageNotice.vue";
import OpenData from "@/components/OpenData/OpenData.vue";
import TagsView from "./components/TagsView.vue";

const formRef = ref(null);
const layoutStore = useLayoutStore();
const appStore = useAppStore();
const userStore = useUserStore();
const router = useRouter();
const { isFullscreen, enter, exit } = useFullscreen();
const userDropOptions = [
  {
    label: "锁定屏幕",
    key: "lockScreen",
    icon: () =>
      h(Icon, { icon: "ant-design:lock-outlined", width: 18, height: 18 }),
  },
  {
    label: "退出登录",
    key: "logout",
    icon: () =>
      h(Icon, { icon: "icon-park-outline:power", width: 16, height: 16 }),
  },
];

let lockScreenModalShow = ref(false);
const dialogFormRef = reactive({
  password: "",
});
const rulesRef = reactive({
  password: {
    required: true,
    message: "请输入锁屏密码",
  },
});

const collapsed = computed({
  get() {
    return layoutStore.collapsed;
  },
  set(isCollapsed) {
    layoutStore.$patch((state) => {
      state.collapsed = isCollapsed;
    });
  },
});
const isLockScreen = computed(() => appStore.lockScreen);

watch(
  () => lockScreenModalShow,
  (isShow) => {
    clearValidate();
    if (!isShow) {
      dialogFormRef.password = "";
    }
  }
);

// 点击用户下拉菜单
function handleUserMenuClick(key) {
  if (key === "logout") {
    logout();
  } else if (key === "lockScreen") {
    lockScreenModalShow.value = true;
  }
}

// 退出登录
function logout() {
  Modal.confirm({
    title: "温馨提醒",
    icon: createVNode(Icon, {
      icon: "ant-design:exclamation-circle-outlined",
      width: 24,
      height: 24,
      inline: true,
    }),
    content: "是否确认退出系统?",
    onOk() {
      return new Promise((resolve) => {
        userStore.logout();
        router.push("/login");
        resolve();
      });
    },
    onCancel() {},
  });
}

// 锁定屏幕
function lockScreen() {
  formRef.value?.validate((errors) => {
    if (!errors) {
      lockScreenModalShow.value = false;
      appStore.lockScreenPassword = dialogFormRef.password;
    }
  });
}

// 清除校验信息
function clearValidate() {
  formRef.value?.restoreValidation();
}
</script>

<style scoped lang="scss">
.header {
  width: 100%;
  background: #fff;
  padding: 0 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  line-height: 1;
  border-bottom: 1px solid #d9d9d9;

  .trigger {
    font-size: 20px;
  }

  .icon-item {
    display: inline-block;
    height: 48px;
    line-height: 48px;
    padding: 0 10px;

    &:hover {
      background: #f6f6f6;
    }
  }

  .icon {
    font-size: 20px;
    vertical-align: middle !important;
  }

  .user-area {
    .username {
      margin-left: 10px;
    }
  }
}
</style>
