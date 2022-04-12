<template>
  <div class="bg-white flex py-3 px-4 items-center justify-between">
    <div class="p-4 cursor-pointer icon">
      <svg
        t="1637409203900"
        class="icon-article"
        viewBox="0 0 1024 1024"
        version="1.1"
        xmlns="http://www.w3.org/2000/svg"
        p-id="10795"
        width="200"
        height="200"
      >
        <path
          d="M807.964324 746.876984a53.965697 53.965697 0 0 0 26.239567-46.22174V161.758841a53.896555 53.896555 0 0 0-53.896555-53.896555h-269.586486a53.931126 53.931126 0 0 0-45.703172 25.340715h168.673053a174.411878 174.411878 0 0 1 174.377307 174.411878zM351.589457 133.133859a161.689664 161.689664 0 0 1 159.027679-133.133824h269.482773a161.689664 161.689664 0 0 1 161.793377 161.758806v538.965545a161.724235 161.724235 0 0 1-134.205532 159.373391 174.342736 174.342736 0 0 1-173.997023 163.902223H256.44942A174.342736 174.342736 0 0 1 82.106685 849.726407V307.476594a174.342736 174.342736 0 0 1 174.342735-174.273593h95.140037zM256.44942 240.926968a66.584198 66.584198 0 0 0-66.549626 66.618769v542.18067a66.549626 66.549626 0 0 0 66.549626 66.549626h377.241311a66.549626 66.549626 0 0 0 66.480484-66.549626V307.476594a66.549626 66.549626 0 0 0-66.480484-66.515055H256.44942z m52.271707 126.807288h248.912888a53.896555 53.896555 0 1 1 0 107.793109h-248.912888a53.896555 53.896555 0 0 1 0-107.793109z m0 240.926933h248.912888a53.896555 53.896555 0 1 1 0 107.758538h-248.912888a53.896555 53.896555 0 1 1 0-107.758538z m0 0"
          p-id="10796"
        />
      </svg>
    </div>
    <div>
      <div class="text-lg text-gray-400 font-bold">文章数</div>
      <div class="text-lg font-bold">12</div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "ArticleCount",
});
</script>
<script setup></script>

<style scoped lang="scss">
.icon {
  transition: all 0.38s ease-out;
  border-radius: 10px;

  &:hover {
    background-color: #4974ef;

    .icon-article {
      fill: white;
    }
  }

  .icon-article {
    width: 48px;
    height: 48px;
    fill: #4974ef;
  }
}
</style>
