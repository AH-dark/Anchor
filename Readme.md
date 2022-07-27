# Anchor

Self-built proxy system for GitHub repositories, NPM packages, and WordPress themes and plugins.

During initialization, if there is no `config.yaml` file in the current directory, it will be automatically generated.

The content of the initial file is as follows:

```yaml
# 系统设置
system:
    name: Anchor # string, 程序名称
    listen: :8080 # string, gin 监听端点
    debug: false # bool, 是否开启Debug模式
    
# 代理设置
proxy:
    github: # GitHub 代理
        open: false # 是否开启
        minify: onlyMin # 压缩文件配置，共三种：all / onlyMin / none
        endpoint: ['https://raw.githubusercontent.com/{{user}}/{{repo}}/{{version}}/{{path}}'] # array<string> 代理列表，自上而下依次尝试
        white_list: [] # 白名单，格式：<user>/<repo>，支持 * 通配符
    npm:
        open: false # 是否开启
        minify: onlyMin # 压缩文件配置，共三种：all / onlyMin / none
        endpoint: ['https://unpkg.com/{{package}}@{{version}}/{{path}}'] # array<string> 代理列表，自上而下依次尝试
        white_list: [] # 白名单，格式：@<user>/<package> 或 <package>
        
# 跨域配置
cors:
    allow_origins: [UNSET] # 默认为通配符
    allow_methods: [GET, HEAD, OPTIONS]
    allow_headers: [Content-Length, Content-Type, X-Powered-By, X-Run-By, X-Timestamp]
    allow_credentials: false
    expose_headers: []
```
