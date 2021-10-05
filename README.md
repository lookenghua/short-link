# 短链接项目开发实战

## 目录结构
```
short-link
├── config # 配置文件
├── controller #控制器
│   └── user.go
├── ent #数据库模型
│   ├── enttest
│   │   └── enttest.go
│   ├── hook
│   │   └── hook.go
│   ├── migrate
│   │   ├── migrate.go
│   │   └── schema.go
│   ├── predicate
│   │   └── predicate.go
│   ├── runtime
│   │   └── runtime.go
│   ├── schema
│   │   └── user.go
│   ├── user
│   │   ├── user.go
│   │   └── where.go
│   ├── client.go
│   ├── config.go
│   ├── context.go
│   ├── ent.go
│   ├── generate.go
│   ├── mutation.go
│   ├── runtime.go
│   ├── tx.go
│   ├── user.go
│   ├── user_create.go
│   ├── user_delete.go
│   ├── user_query.go
│   └── user_update.go
├── repository # 数据库操作
│   └── user.go
├── router # 路由
│   ├── middleware #中间件
│   └── router.go #路由配置
├── service # 服务层
├── util # 工具集
├── bee.json
├── go.mod
├── go.sum
├── main.go
├── README.md
```