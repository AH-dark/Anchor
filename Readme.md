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
    wordpress: # WordPress proxy
      plugin_open: false # Whether to open for plugins proxy
      theme_open: false # Whether to open for plugins proxy
      minify: none # Compressed file configuration, only support two types: all / none
      plugin_white_list: [] # Whitelist for plugins, format: <name>
      theme_white_list: [] # Whitelist for themes, format: <name>

# Cross Origin Configuration
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

## Use

### GitHub

The GitHub proxy service obtains resources from https://raw.githubusercontent.com by default in the following format:

```
https://<domain>/gh/<user>/<repo>/<version>/<path>
```

### NPM

The NPM proxy service obtains resources from https://unpkg.com by default in the following format:

```
### Full
https://<domain>/npm/@<user>/<package>@<version>/<path>

### No version
https://<domain>/npm/@<user>/<package>/<path>
# equal to
https://<domain>/npm/@<user>/<package>@latest/<path>

### Standard package
https://<domain>/npm/<package>@<version>/<path>

### No path
https://<domain>/npm/<package>@<version>
# will be redirected to
https://<domain>/npm/@<user>/<package>@<version>/<path> # main file
```

### WordPress

The WordPress proxy service obtains resources from https://themes.svn.wordpress.org and https://plugins.svn.wordpress.org by default in the following format:

```
# For themes in wordpress.org
https://<domain>/wp/theme/<name>/<version>/<path>

# For plugins in wordpress.org
https://<domain>/wp/plugin/<name>/<version>/<path>
```

## Feature

### Auto compress

The automatic compression of each service is determined according to the `config.yaml` filled in by the user.

It can be used only if the `minify` item is set to `all` or `onlyMin` in the configuration file of the corresponding service.

We will explain to you in detail the difference between the three modes.

#### Support

The currently supported file types are as follows:

* text/html
* application/javascript
* text/css
* application/json

[tdewolff/minify](https://github.com/tdewolff/minify) provides file compression support for this project.

#### Compress

The compression mode of different applications is different, you need to configure the `minify` option of the corresponding application in `config.yaml`.

Our automatic compression follows these rules:

* Priority back to source files with the origin path.
* When the file with the origin path is not available, it will try to remove the `.min.` part that exists in the path and try to get it again.

##### All

When you set `minify` to `all` in `config.yaml`, the system will compress **all** files that can be compressed.

##### OnlyMin

In this mode, the system compresses **only files ending in `.min.*`**, such as `.min.js`, `.min.css`, `.min.json`.

##### None

In this mode, the system will **not compress any files**.
