# goweibo

Go Weibo App

- v1 版本地址: https://github.com/Away0x/goweibo/tree/master
  - 前后端未分离，技术栈为:
    - Web framework: [gin](https://github.com/gin-gonic/gin)
    - ORM: [gorm](https://github.com/go-gorm/gorm)
    - 热重载: [fresh](https://github.com/Away0x/fresh)
- v2 版本地址: https://github.com/Away0x/goweibo/tree/v2
  - 前后端分离:
    - 后端技术栈:
      - Web framework: [echo](https://github.com/labstack/echo)
      - ORM: [gorm](https://github.com/go-gorm/gorm)
      - 日志: [zap](https://github.com/uber-go/zap)
      - 参数验证: [govalidator](https://github.com/thedevsaddam/govalidator)
      - 配置: [viper](github.com/spf13/viper)
      - 热重载: [air](https://github.com/cosmtrek/air)
      - 命令行: [cobra](https://github.com/spf13/cobra)
      - API 文档: [swagger](https://github.com/go-swagger/go-swagger)
    - 前端技术栈:
      - typescript、react、react hooks
      - 测试: jest, enzyme
      - UI 框架: [ant-design](https://github.com/ant-design/ant-design)
      - CSS: [sass](https://github.com/sass/node-sass), [styled-components](https://github.com/styled-components/styled-components)
      - 状态管理: [unstated-next](https://github.com/jamiebuilds/unstated-next), [immer]()
      - Event emitter: [mitt](https://github.com/developit/mitt)
      - Hooks Library: [hooks](https://github.com/alibaba/hooks)
      - HTTP Client: [axios](https://github.com/axios/axios)
      - UI 文档: [storybook](https://github.com/storybookjs/storybook)
      - 代码生成工具: [hygen](https://github.com/jondot/hygen)

## 后端项目结构

## 前端
### 启动项目

### 部署

### Npm Scripts

### 项目结构目录
<details>
<summary>展开查看</summary>
<pre><code>
├── assets           图片字体等资源
│
├── components       公用组件
│
├── config           配置
│
├── constants        常量
│
├── containers       状态容器
│
├── events           事件 (通常用于 view 和 services/tools 的解耦)
│
├── layouts          布局
│
├── pages            页面
│
├── routes           路由
│
├── services         数据层 (网络数据/本地存储数据/mock 数据)
│
├── styles           样式
│
├── tools            工具
│
├── typings          类型定义
│
├── App.tsx          根组件
│
└── index.tsx        入口
</code></pre>
</details>

### 其他
#### vscode
##### 扩展
- EditorConfig
- Prettier
- ESLint
- vscode-styled-components
