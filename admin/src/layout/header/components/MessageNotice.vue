<template>
  <a-popover placement="bottom">
    <template #content>
      <div style="width: 250px">
        <a-tabs v-model:activeKey="activeKey">
          <a-tab-pane key="1" tab="消息"> Content of Tab Pane 1</a-tab-pane>
          <a-tab-pane key="2" tab="评论"> </a-tab-pane>
          <a-tab-pane key="3" tab="其他"> Content of Tab Pane 3</a-tab-pane>
        </a-tabs>
      </div>
    </template>
    <a-badge :count="todoCount" show-zero>
      <Icon
        icon="ant-design:bell-outlined"
        width="24"
        height="24"
        :inline="true"
        class="icon"
      />
    </a-badge>
  </a-popover>
</template>

<script setup>
import { ref } from "vue";

let activeKey = ref("2");
const todoCount = ref(0);
</script>

<style scoped>
.icon {
  font-size: 20px;
  vertical-align: middle !important;
}
</style>
