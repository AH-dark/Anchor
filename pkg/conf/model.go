package conf

// system 系统设置
type system struct {
	Name   string `yaml:"name"`
	Listen string `yaml:"listen"`
	Debug  bool   `yaml:"debug"`
}

// githubProxy GitHub镜像反向代理配置
type githubProxy struct {
	Open      bool       `yaml:"open"`
	Minify    MinifyType `yaml:"minify"`
	Endpoint  []string   `yaml:"endpoint,flow"`
	WhiteList []string   `yaml:"white_list,flow"`
}

// npmProxy Npm镜像反向代理配置
type npmProxy struct {
	Open      bool       `yaml:"open"`
	Minify    MinifyType `yaml:"minify"`
	Endpoint  []string   `yaml:"endpoint,flow"`
	WhiteList []string   `yaml:"white_list,flow"`
}

type wordpressProxy struct {
	PluginOpen      bool       `yaml:"plugin_open"`
	ThemeOpen       bool       `yaml:"theme_open"`
	Minify          MinifyType `yaml:"minify"`
	PluginWhiteList []string   `yaml:"plugin_white_list,flow"`
	ThemeWhiteList  []string   `yaml:"theme_white_list,flow"`
}

// proxy 反向代理配置
type proxy struct {
	Github githubProxy    `yaml:"github"`
	Npm    npmProxy       `yaml:"npm"`
	Wp     wordpressProxy `yaml:"wordpress"`
}

// cors 跨域配置
type cors struct {
	AllowOrigins     []string `yaml:"allow_origins,flow"`
	AllowMethods     []string `yaml:"allow_methods,flow"`
	AllowHeaders     []string `yaml:"allow_headers,flow"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	ExposeHeaders    []string `yaml:"expose_headers,flow"`
}

// config 配置
type config struct {
	System system `yaml:"system"`
	Proxy  proxy  `yaml:"proxy"`
	CORS   cors   `yaml:"cors"`
}
