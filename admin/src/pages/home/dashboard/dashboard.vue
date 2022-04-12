<template>
  <div class="dashboard-container">
    <div class="h-full">
      <div class="grid grid-cols-4 gap-x-10">
        <ArticleCount />
        <VisitedCount />
        <TopicCount />
        <CommentCount />
      </div>
      <UserVisitedLine class="mt-4 bg-white" />
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "DashboardPage",
});
</script>
<script setup>
import ArticleCount from "./components/ArticleCount.vue";
import CommentCount from "@/pages/home/dashboard/components/CommentCount.vue";
import TopicCount from "@/pages/home/dashboard/components/TopicCount.vue";
import UserVisitedLine from "@/pages/home/dashboard/components/UserVisitedLine.vue";
import VisitedCount from "./components/VisitedCount.vue";
</script>

<style scoped lang="scss">
.dashboard-container {
  padding: 20px;
  background: #f0f2f5;
}
</style>
