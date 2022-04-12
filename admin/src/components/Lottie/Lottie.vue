<template>
  <div ref="boxRef" class="box"></div>
</template>

<script setup>
  import { onMounted, ref } from 'vue'
  import lottie from 'lottie-web'

  const props = defineProps({
    animationData: {
      type: Object,
      required: true,
    },
    loop: {
      type: Boolean,
      default: true,
    },
    autoplay: {
      type: Boolean,
      default: true,
    },
  })
  let boxRef = ref(null)
  onMounted(() => {
    lottie.loadAnimation({
      container: boxRef.value, // the dom element that will contain the animation
      renderer: 'svg',
      loop: props.loop,
      autoplay: props.autoplay,
      animationData: props.animationData,
    })
  })
</script>

<style scoped></style>
