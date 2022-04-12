import { createApp, h } from 'vue'

export function createApi(component) {
  // 显示
  function show(config = {}) {
    let componentNode = document.createElement('div')
    let attr = document.createAttribute('class')
    attr.value = 'create-' + (component.name || component.displayName || component.componentName)
    componentNode.setAttributeNode(attr)
    let app, Comp
    Comp = () => h(component, { ...config })
    app = createApp(Comp)
    app.mount(componentNode)
    document.body.appendChild(componentNode)
    window.addEventListener('hashchange', () => {
      remove()
    })
    window.addEventListener('popstate', () => {
      remove()
    })

    const remove = () => {
      app.unmount(componentNode)
    }
    return {
      remove,
    }
  }

  return {
    show,
  }
}
