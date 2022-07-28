# Anchor

[简体中文](Readme.zh.md)

Self-built proxy system for GitHub repositories, NPM packages, and WordPress themes and plugins.

During initialization, if there is no `config.yaml` file in the current directory, it will be automatically generated.

## Deploy

Currently, we only provide packaged applications, including support for four system architectures.

1. Go to <https://github.com/AH-dark/Anchor/releases>, get the latest release.
2. Download the file and unzip it to get the executable within.
3. Run the executable file, the `config.yaml` file will be automatically generated during initialization, and you can configure it by yourself.

### Configuration

```yaml
# System settings
system:
    name: Anchor # string, application name
    listen: :8080 # string, Gin listen endpoint
    debug: false # bool, whether to enable debug mode
    
# Proxy settings
proxy:
    github: # GitHub proxy
        open: false # Whether to open
        minify: onlyMin # Compressed file configuration, three types: all / onlyMin / none
        endpoint: ['https://raw.githubusercontent.com/{{user}}/{{repo}}/{{version}}/{{path}}'] #array<string> list of proxies, tried in order from top to bottom
        white_list: [] # Whitelist, format: <user>/<repo>, supports * wildcard
    npm: # NPM proxy
        open: false # Whether to open
        minify: onlyMin # Compressed file configuration, three types: all / onlyMin / none
        endpoint: ['https://unpkg.com/{{package}}@{{version}}/{{path}}'] # array<string> list of proxies, tried in order from top to bottom
        white_list: [] # Whitelist, format: @<user>/<package> or <package>
        
# 跨域配置
cors:
    allow_origins: [UNSET] # Default is wildcard
    allow_methods: [GET, HEAD, OPTIONS]
    allow_headers: [Content-Length, Content-Type, X-Powered-By, X-Run-By, X-Timestamp]
    allow_credentials: false
    expose_headers: []
```

### Port

Gin will listen according to the port information in `config.yaml`, which is `0.0.0.0:8080` by default in the package.

Note that Gin omits the `0.0.0.0` prefix by default, but you need to add `:` before the port number.
