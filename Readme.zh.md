# Anchor

[English](Readme.md)

GitHub 仓库、NPM 包、WordPress 主题和插件的自建代理系统。

初始化时，如果当前目录下没有`config.yaml`文件，会自动生成。

## 部署

目前，我们只提供打包的应用程序，包括对四种系统架构的支持。

1. 前往<https://github.com/AH-dark/Anchor/releases>，获取最新版本。
2. 下载文件并解压缩以获取其中的可执行文件。
   3.运行可执行文件，初始化时会自动生成`config.yaml`文件，你可以自己配置。

### 配置

```yaml
# 系统设置
system:
    name: Anchor # 字符串，应用程序名称
    listen: :8080 # 字符串，Gin 监听端点
    debug: false # bool，是否开启调试模式
    
# 代理设置
proxy:
    github: # GitHub 代理
        open: false # 是否打开
        minify: onlyMin # 压缩文件配置，共有三种类型：all / onlyMin / none
        endpoint: ['https://raw.githubusercontent.com/{{user}}/{{repo}}/{{version}}/{{path}}'] #array<string> 代理列表，从上到下依次尝试
        white_list: [] # 白名单，格式：<user>/<repo>，支持*通配符
    npm: # NPM 代理
        open: false # 是否打开
        minify: onlyMin # 压缩文件配置，共有三种类型：all / onlyMin / none
        endpoint: ['https://unpkg.com/{{package}}@{{version}}/{{path}}'] # array<string> 代理列表，从上到下依次尝试
        white_list: [] # 白名单，格式：@<user>/<package> or <package>
        
# 跨域配置
  cors：
    allow_origins: [UNSET] # 默认为通配符
    allow_methods: [GET, HEAD, OPTIONS]
    allow_headers: [Content-Length, Content-Type, X-Powered-By, X-Run-By, X-Timestamp]
    allow_credentials: false
    expose_headers: []
```

### 端口

Gin 会根据 `config.yaml` 中的端口信息进行监听，软件包中默认为 `0.0.0.0:8080`。

需要注意的是，Gin 默认省略了 `0.0.0.0` 前缀，但您仍需要在端口号前添加 `:` 符号。
