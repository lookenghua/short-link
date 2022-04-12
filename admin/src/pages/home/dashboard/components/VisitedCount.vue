<template>
  <div class="bg-white flex py-3 px-4 items-center justify-between">
    <div class="p-4 cursor-pointer icon">
      <svg
        t="1637414616376"
        class="icon-visited"
        viewBox="0 0 1024 1024"
        version="1.1"
        xmlns="http://www.w3.org/2000/svg"
        p-id="6660"
      >
        <path
          d="M510.557138 204.354248c-302.839283 0-458.323497 306.150702-458.323497 306.150702s117.330242 306.189587 458.323497 306.189587c308.802088 0 458.300984-304.995389 458.300984-304.995389S818.167075 204.354248 510.557138 204.354248L510.557138 204.354248 510.557138 204.354248zM511.245823 701.866279c-110.729917 0-190.772928-83.72589-190.772928-191.364399 0-107.647719 80.049151-191.352119 190.772928-191.352119 110.723777 0 190.763718 83.697237 190.763718 191.352119C702.010565 618.140389 621.970624 701.866279 511.245823 701.866279L511.245823 701.866279 511.245823 701.866279zM511.245823 395.675668c-63.286372 0.145309-114.460892 53.321416-114.460892 114.827235 0 61.473073 51.175543 114.821095 114.460892 114.821095 63.282279 0 114.453728-53.352115 114.453728-114.821095C625.703645 448.975595 574.529125 395.556964 511.245823 395.675668L511.245823 395.675668 511.245823 395.675668z"
          p-id="6661"
        />
      </svg>
    </div>
    <div>
      <div class="text-lg text-gray-400 font-bold">浏览人数</div>
      <div class="text-lg font-bold">12</div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "VisitedCount",
});
</script>
<script setup></script>

<style scoped lang="scss">
.icon {
  transition: all 0.38s ease-out;
  border-radius: 10px;

  &:hover {
    background-color: #f1723b;

    .icon-visited {
      fill: white;
    }
  }

  .icon-visited {
    width: 48px;
    height: 48px;
    fill: #f1723b;
  }
}
</style>
