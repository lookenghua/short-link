export function Flipper(config) {
  // 默认配置
  this.config = {
    // 时钟模块的节点
    node: null,
    // 初始前牌文字
    frontText: 'number0',
    // 初始后牌文字
    backText: 'number1',
    // 翻转动画时间（毫秒，与翻转动画CSS 设置的animation-duration时间要一致）
    duration: 600,
  }
  // 节点的原本class，与html对应，方便后面添加/删除新的class
  this.nodeClass = {
    flip: 'flip',
    front: 'digital front',
    back: 'digital back',
  }
  // 覆盖默认配置
  Object.assign(this.config, config)
  // 定位前后两个牌的DOM节点
  this.frontNode = this.config.node.querySelector('.front')
  this.backNode = this.config.node.querySelector('.back')
  // 是否处于翻牌动画过程中（防止动画未完成就进入下一次翻牌）
  this.isFlipping = false
  // 初始化
  this._init()
}
Flipper.prototype = {
  constructor: Flipper,
  // 初始化
  _init: function () {
    // 设置初始牌面字符
    this._setFront(this.config.frontText)
    this._setBack(this.config.backText)
  },
  // 设置前牌文字
  _setFront: function (className) {
    this.frontNode.setAttribute('class', this.nodeClass.front + ' ' + className)
  },
  // 设置后牌文字
  _setBack: function (className) {
    this.backNode.setAttribute('class', this.nodeClass.back + ' ' + className)
  },
  _flip: function (type, front, back) {
    // 如果处于翻转中，则不执行
    if (this.isFlipping) {
      return false
    }
    // 设置翻转状态为true
    this.isFlipping = true
    // 设置前牌文字
    this._setFront(front)
    // 设置后牌文字
    this._setBack(back)
    // 根据传递过来的type设置翻转方向
    let flipClass = this.nodeClass.flip
    if (type === 'down') {
      flipClass += ' down'
    } else {
      flipClass += ' up'
    }
    // 添加翻转方向和执行动画的class，执行翻转动画
    this.config.node.setAttribute('class', flipClass + ' go')
    // 根据设置的动画时间，在动画结束后，还原class并更新前牌文字
    setTimeout(() => {
      // 还原class
      this.config.node.setAttribute('class', flipClass)
      // 设置翻转状态为false
      this.isFlipping = false
      // 将前牌文字设置为当前新的数字，后牌因为被前牌挡住了，就不用设置了。
      this._setFront(back)
    }, this.config.duration)
  },
  // 下翻牌
  flipDown: function (front, back) {
    this._flip('down', front, back)
  },
  // 上翻牌
  flipUp: function (front, back) {
    this._flip('up', front, back)
  },
}
