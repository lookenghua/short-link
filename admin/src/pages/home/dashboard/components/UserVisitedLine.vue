<template>
  <div ref="chartRef" style="height: 366px"></div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "UserVisitedLine",
});
</script>
<script setup>
import { init } from "echarts";
import { onMounted, ref } from "vue";

let chartRef = ref(null);
onMounted(() => {
  const chart = init(chartRef.value);
  chart.setOption({
    title: {
      text: "近30分钟访问量",
      left: "center",
      top: 20,
    },
    grid: {
      left: 50,
      right: 80,
      top: 60,
      bottom: 40,
    },
    tooltip: {},
    xAxis: {
      name: "访问时间",
      data: ["衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"],
      axisTick: {
        show: true,
        alignWithLabel: true,
      },
    },
    yAxis: {
      name: "访问人次",
      axisLine: {
        show: true,
      },
    },
    series: [
      {
        name: "访问量",
        type: "line",
        data: [5, 20, 36, 10, 10, 20],
        smooth: true,
        label: {
          show: true,
        },
      },
    ],
  });
  window.addEventListener("resize", () => {
    chart.resize();
  });
});
</script>

<style scoped></style>
