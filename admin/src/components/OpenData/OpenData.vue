<template>
  <n-avatar
    round
    v-if="type === 'avatar'"
    :class="props.class"
    :src="avatar"
    :size="props.size"
    :alt="username[0]"
  ></n-avatar>
  <span v-if="type === 'username'" :class="props.class">{{ username }}</span>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "OpenData",
});
</script>
<script setup>
import { completeLink } from "@/utils/file.js";
import { computed, toRefs } from "vue";
import { useUserStore } from "@/store/modules/user.js";

const props = defineProps({
  type: {
    type: String,
    required: true,
  },
  size: {
    type: Number,
    default: 32,
  },
  class: {
    type: String,
    default: "",
  },
});
const { type } = toRefs(props);
const userStore = useUserStore();
const userInfo = computed(() => userStore.userInfo);
const avatar = computed(() => completeLink(userInfo.value?.avatar));
const username = computed(() => userInfo.value?.username || "Admin");
</script>

<style scoped></style>
