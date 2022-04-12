<template>
  <a-layout-header class="header-area z-40">
    <div class="header">
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
      <div>
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
          <a-dropdown>
            <div class="user-area">
              <OpenData type="avatar" />
              <OpenData type="username" class="text-base" />
            </div>
            <template #overlay>
              <a-menu @click="handleUserMenuClick">
                <a-menu-item key="lockScreen">
                  <div class="flex items-center">
                    <Icon
                      icon="ant-design:lock-outlined"
                      width="18"
                      height="18"
                      :inline="true"
                    />
                    <span class="ml-1">锁定屏幕</span>
                  </div>
                </a-menu-item>
                <a-menu-item key="logout">
                  <div class="flex items-center">
                    <Icon
                      icon="icon-park-outline:power"
                      width="16"
                      height="16"
                      :inline="true"
                    />
                    <span class="ml-1">退出系统</span>
                  </div>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>
      </div>
    </div>
    <TagsView />
  </a-layout-header>
  <a-modal
    v-model:visible="lockScreenModalShow"
    title="锁定屏幕"
    :footer="null"
  >
    <div class="pb-8">
      <div class="flex flex-col items-center">
        <OpenData type="avatar" :size="64" class="align-text-bottom" />
        <OpenData type="username" class="text-lg" />
      </div>
      <div class="text-base mb-2 mt-6">
        <span class="text-red-500">*</span><span>锁屏密码</span>
      </div>
      <a-form>
        <a-form-item v-bind="validateInfos.password">
          <a-input-password
            v-model:value="dialogFormRef.password"
            placeholder="请输入锁屏密码"
            @focus="clearValidate()"
          />
        </a-form-item>
      </a-form>
      <a-button
        type="primary"
        block
        class="mt-3 !rounded-sm"
        @click="lockScreen"
        >锁定</a-button
      >
    </div>
  </a-modal>
  <LockScreen v-if="isLockScreen" />
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({ name: "LayoutHeader" });
</script>
<script setup>
import { computed, createVNode, reactive, ref, watch } from "vue";
import { Form } from "ant-design-vue";
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

const useForm = Form.useForm;
const layoutStore = useLayoutStore();
const appStore = useAppStore();
const userStore = useUserStore();
const router = useRouter();
const { isFullscreen, enter, exit } = useFullscreen();

let lockScreenModalShow = ref(false);
const dialogFormRef = reactive({
  password: "",
});
const rulesRef = reactive({
  password: [
    {
      required: true,
      message: "请输入锁屏密码",
    },
  ],
});
const { clearValidate, validate, validateInfos, resetFields } = useForm(
  dialogFormRef,
  rulesRef
);

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
const leftDistance = computed(() => (collapsed.value ? "48px" : "200px"));
const isLockScreen = computed(() => appStore.lockScreen);

watch(
  () => lockScreenModalShow,
  (isShow) => {
    if (!isShow) {
      dialogFormRef.password = "";
      resetFields();
    }
  }
);

// 点击用户下拉菜单
function handleUserMenuClick({ key }) {
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
  validate()
    .then(() => {
      lockScreenModalShow.value = false;
      appStore.lockScreenPassword = dialogFormRef.password;
    })
    .catch((err) => {
      console.log(err);
    });
}
</script>

<style scoped lang="scss">
.header-area {
  height: auto;
  padding: 0;
  position: fixed;
  top: 0;
  left: v-bind(leftDistance);
  right: 0;

  transition: all 0.2s;

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
}
</style>
