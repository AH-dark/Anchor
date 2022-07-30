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
    wordpress: # WordPress 代理
        plugin_open: false # 是否开启插件代理
        theme_open: false # 是否开启插件代理
        minify: none # 压缩文件配置，只支持两种：all / none
        plugin_white_list: [] # 插件白名单，格式：<name>
        theme_white_list: [] # 主题白名单，格式：<name>

# 跨域配置
cors:
    allow_origins: [UNSET] # 默认为通配符
    allow_methods: [GET, HEAD, OPTIONS]
    allow_headers: [Content-Length, Content-Type, X-Powered-By, X-Run-By, X-Timestamp]
    allow_credentials: false
    expose_headers: []
```

### 端口

Gin 会根据 `config.yaml` 中的端口信息进行监听，软件包中默认为 `0.0.0.0:8080`。

需要注意的是，Gin 默认省略了 `0.0.0.0` 前缀，但您仍需要在端口号前添加 `:` 符号。

## 使用

### GitHub

GitHub代理服务默认从 https://raw.githubusercontent.com 获取资源，格式如下：

```
https://<域>/gh/<用户>/<repo>/<版本>/<路径>
```

### NPM

NPM代理服务默认从 https://unpkg.com 获取资源，格式如下：

```
### 完整的
https://<域名>/npm/@<组织名>/<包名>@<版本>/<路径>

### 无版本标识
https://<域名>/npm/@<组织名>/<包名>/<路径>
# 等于
https://<域名>/npm/@<组织名>/<包名>@latest/<路径>

### 标准的
https://<域名>/npm/<包名>@<版本>/<路径>

### 无路径的
https://<domain>/npm/<package>@<version>
# 将被重定向到
https://<domain>/npm/@<user>/<package>@<version>/<path> # main 文件（来自包内 package.json 定义）
```

### WordPress

WordPress 代理服务默认从 https://themes.svn.wordpress.org 和 https://plugins.svn.wordpress.org 获取资源，格式如下：

```
# 对于 wordpress.org 中的上架主题
https://<domain>/wp/theme/<name>/<version>/<path>

# 对于 wordpress.org 中的上架插件
https://<domain>/wp/plugin/<name>/<version>/<path>
```

## 特性

### 自动压缩

每个服务的自动压缩是根据用户填写的`config.yaml`来决定的。

只有在相应服务的配置文件中将`minify`项设置为`all`或`onlyMin`时才能使用。

我们将为您详细解释这三种模式之间的区别。

#### 支持

目前支持的文件类型如下：

* 文本/html
* 应用程序/javascript
* 文本/css
* 应用程序/json

[tdewolff/minify](https://github.com/tdewolff/minify)为本项目提供文件压缩支持。

#### 压缩

不同应用的压缩方式不同，需要在`config.yaml`中配置对应应用的`minify`选项。

我们的自动压缩始终遵循以下规则：

* 优先返回具有原始路径的源文件。
* 当原始路径的文件不可用时，它会尝试删除路径中存在的`.min.`部分，并尝试重新获取。

##### All

在 `config.yaml` 中将 `minify` 设置为 `all` 时，系统会**压缩全部可以压缩的文件**。

##### OnlyMin

在此模式下，系统**仅压缩以`.min.*`结尾的文件**，例如`.min.js`、`.min.css`、`.min.json`。

##### None

在这种模式下，系统将**不压缩任何文件**。