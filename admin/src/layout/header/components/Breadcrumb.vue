<template>
  <n-breadcrumb>
    <n-breadcrumb-item v-for="(item, index) in routes" :key="index">
      {{ item }}
    </n-breadcrumb-item>
  </n-breadcrumb>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "BreadcrumbList",
});
</script>
<script setup>
import { ref, watch } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
let routes = ref([]);
watch(
  route,
  () => {
    routes.value = route.matched.map((it) => it?.meta?.title);
  },
  { immediate: true }
);
</script>

<style scoped></style>
