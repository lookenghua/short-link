<template>
  <div class="bg-white flex py-3 px-4 items-center justify-between">
    <div class="p-4 cursor-pointer icon">
      <svg
        t="1637414852203"
        class="icon-comment"
        viewBox="0 0 1024 1024"
        version="1.1"
        xmlns="http://www.w3.org/2000/svg"
        p-id="20378"
      >
        <path
          d="M512 0C226.742857 0 0 197.485714 0 446.171429c0 138.971429 73.142857 270.628571 190.171429 351.085714L190.171429 1024l226.742857-138.971429c29.257143 7.314286 65.828571 7.314286 95.085714 7.314286 285.257143 0 512-197.485714 512-446.171429C1024 197.485714 797.257143 0 512 0zM256 512C219.428571 512 190.171429 482.742857 190.171429 446.171429S219.428571 380.342857 256 380.342857c36.571429 0 65.828571 29.257143 65.828571 65.828571S292.571429 512 256 512zM512 512C475.428571 512 446.171429 482.742857 446.171429 446.171429S475.428571 380.342857 512 380.342857c36.571429 0 65.828571 29.257143 65.828571 65.828571S548.571429 512 512 512zM768 512C731.428571 512 702.171429 482.742857 702.171429 446.171429s29.257143-65.828571 65.828571-65.828571c36.571429 0 65.828571 29.257143 65.828571 65.828571S804.571429 512 768 512z"
          p-id="20379"
        />
      </svg>
    </div>
    <div>
      <div class="text-lg text-gray-400 font-bold">评论数</div>
      <div class="text-lg font-bold">12</div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "CommentCount",
});
</script>
<script setup></script>

<style scoped lang="scss">
.icon {
  transition: all 0.38s ease-out;
  border-radius: 10px;

  &:hover {
    background-color: #1296db;

    .icon-comment {
      fill: white;
    }
  }

  .icon-comment {
    width: 48px;
    height: 48px;
    fill: #1296db;
  }
}
</style>
