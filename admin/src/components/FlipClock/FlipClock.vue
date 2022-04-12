<template>
  <div id="clock" ref="clock" class="clock">
    <div class="flip down">
      <div class="digital front number0"></div>
      <div class="digital back number1"></div>
    </div>
    <div class="flip down">
      <div class="digital front number0"></div>
      <div class="digital back number1"></div>
    </div>
    <div class="divider">:</div>
    <div class="flip down">
      <div class="digital front number0"></div>
      <div class="digital back number1"></div>
    </div>
    <div class="flip down">
      <div class="digital front number0"></div>
      <div class="digital back number1"></div>
    </div>
    <div class="divider">:</div>
    <div class="flip down">
      <div class="digital front number0"></div>
      <div class="digital back number1"></div>
    </div>
    <div class="flip down">
      <div class="digital front number0"></div>
      <div class="digital back number1"></div>
    </div>
  </div>
</template>
<script>
  import { defineComponent } from 'vue'

  export default defineComponent({
    name: 'FlipClock',
  })
</script>
<script setup>
  // 定位时钟模块
  import { Flipper } from '@/components/FlipClock/flipClock.js'
  import { onMounted, ref, watch } from 'vue'
  import { useWindowSize } from '@vueuse/core'

  let clock = ref(null)
  let size = ref(200)
  const { width } = useWindowSize()

  watch(width, () => {
    size.value = (width.value / 1920) * 200
  })

  //正则格式化日期
  function formatDate(date, dateFormat) {
    /* 单独格式化年份，根据y的字符数量输出年份
* 例如：yyyy => 2019
yy => 19
y => 9
*/
    if (/(y+)/.test(dateFormat)) {
      dateFormat = dateFormat.replace(
        RegExp.$1,
        (date.getFullYear() + '').substr(4 - RegExp.$1.length)
      )
    }
    // 格式化月、日、时、分、秒
    let o = {
      'm+': date.getMonth() + 1,
      'd+': date.getDate(),
      'h+': date.getHours(),
      'i+': date.getMinutes(),
      's+': date.getSeconds(),
    }
    for (let k in o) {
      if (new RegExp(`(${k})`).test(dateFormat)) {
        // 取出对应的值
        let str = o[k] + ''
        /* 根据设置的格式，输出对应的字符
         * 例如: 早上8时，hh => 08，h => 8
         * 但是，当数字>=10时，无论格式为一位还是多位，不做截取，这是与年份格式化不一致的地方
         * 例如: 下午15时，hh => 15, h => 15
         */
        dateFormat = dateFormat.replace(RegExp.$1, RegExp.$1.length === 1 ? str : padLeftZero(str))
      }
    }
    return dateFormat
  }

  //日期时间补零
  function padLeftZero(str) {
    return ('00' + str).substr(str.length)
  }

  onMounted(() => {
    size.value = (clock.value.offsetWidth / 1920) * 200
    // 定位6个翻板
    let flips = clock.value.querySelectorAll('.flip')
    // 获取当前时间
    let now = new Date()
    // 格式化当前时间，例如现在是20:30:10，则输出"203010"字符串
    let nowTimeStr = formatDate(now, 'hhiiss')
    // 格式化下一秒的时间
    let nextTimeStr = formatDate(new Date(now.getTime() + 1000), 'hhiiss')
    // 定义牌板数组，用来存储6个Flipper翻板对象
    let flipObjs = []
    for (let i = 0; i < flips.length; i++) {
      // 创建6个Flipper实例，并初始化
      flipObjs.push(
        new Flipper({
          // 每个flipper实例按数组顺序与翻板DOM的顺序一一对应
          node: flips[i],
          // 按数组顺序取时间字符串对应位置的数字
          frontText: 'number' + nowTimeStr[i],
          backText: 'number' + nextTimeStr[i],
        })
      )
    }

    // 开始计时
    setInterval(function () {
      // 获取当前时间
      let now = new Date()
      let nowTimeStr = formatDate(new Date(now.getTime() - 1000), 'hhiiss')
      let nextTimeStr = formatDate(now, 'hhiiss')
      for (let i = 0; i < flipObjs.length; i++) {
        if (nowTimeStr[i] === nextTimeStr[i]) {
          continue
        }
        flipObjs[i].flipDown('number' + nowTimeStr[i], 'number' + nextTimeStr[i])
      }
    }, 1000)
  })
</script>

<style scoped lang="scss">
  $bg: #dddddd;
  $color: #151515;

  @font-face {
    font-family: 'DIN Bold';
    src: url('~@/assets/font/din-bold-2.ttf');
  }

  .clock {
    width: 100%;
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;

    .divider {
      display: inline-block;
      height: auto;
      font-size: calc(v-bind(size) * 66 / 60 * 1px);
      font-style: normal;
      color: #dddddd;
      margin-right: calc(v-bind(size) * 10 / 60 * 1px);
    }
  }

  .flip {
    display: inline-block;
    position: relative;
    width: calc(v-bind(size) * 1px);
    height: calc(v-bind(size) * 5 / 3 * 1px);
    line-height: calc(v-bind(size) * 5 / 3 * 1px);
    border-radius: calc(v-bind(size) * 10 / 60 * 1px);
    font-size: calc(v-bind(size) * 66 / 60 * 1px);
    text-align: center;
    font-family: 'DIN Bold';
    font-weight: bold;
    margin-right: calc(v-bind(size) * 10 / 60 * 1px);

    &:before {
      content: '';
      display: block;
      position: absolute;
      top: calc((100% - 1px) / 2);
      left: 0;
      right: 0;
      height: 1px;
      background: #666;
      z-index: 4;
    }

    .digital:before,
    .digital:after {
      content: '';
      position: absolute;
      left: 0;
      right: 0;
      overflow: hidden;
      box-sizing: border-box;
    }

    .digital:before {
      top: 0;
      bottom: 50%;
      border-radius: 10px 10px 0 0;
      //border-bottom: solid 1px #666;
      color: #222222;
      background-color: #dddddd;
    }

    .digital:after {
      top: 50%;
      bottom: 0;
      border-radius: 0 0 10px 10px;
      line-height: 0;
      color: #333333;
      background-color: #eeeeee;
    }
  }

  /*向下翻*/
  .flip.down {
    .front:before {
      z-index: 3;
    }

    .back:after {
      z-index: 2;
      transform-origin: 50% 0;
      transform: perspective(160px) rotateX(180deg);
    }

    .front:after,
    .back:before {
      z-index: 1;
    }
  }

  .flip.down.go {
    .front:before {
      transform-origin: 50% 100%;
      animation: frontFlipDown 0.6s ease-in-out both;
      box-shadow: 0 -2px 6px rgba(255, 255, 255, 0.3);
      backface-visibility: hidden;
    }

    .back:after {
      animation: backFlipDown 0.6s ease-in-out both;
    }
  }

  /*向上翻*/
  .flip.up {
    .front:after {
      z-index: 3;
    }

    .back:before {
      z-index: 2;
      transform-origin: 50% 100%;
      transform: perspective(160px) rotateX(-180deg);
    }

    .front:before,
    .back:after {
      z-index: 1;
    }
  }

  .flip.up.go {
    .front:after {
      transform-origin: 50% 0;
      animation: frontFlipUp 0.6s ease-in-out both;
      box-shadow: 0 2px 6px rgba(255, 255, 255, 0.3);
      backface-visibility: hidden;
    }

    .back:before {
      animation: backFlipUp 0.6s ease-in-out both;
    }
  }

  @keyframes frontFlipDown {
    0% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(0deg);
    }

    100% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(-180deg);
    }
  }

  @keyframes backFlipDown {
    0% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(180deg);
    }

    100% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(0deg);
    }
  }

  @keyframes frontFlipUp {
    0% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(0deg);
    }

    100% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(180deg);
    }
  }

  @keyframes backFlipUp {
    0% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(-180deg);
    }

    100% {
      transform: perspective(calc(v-bind(size) * 160 / 60 * 1px)) rotateX(0deg);
    }
  }

  $numbers: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9;

  @each $num in $numbers {
    .flip .number#{$num}:before,
    .flip .number#{$num}:after {
      content: '#{$num}';
    }
  }
</style>
