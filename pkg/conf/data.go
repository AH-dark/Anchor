package conf

type MinifyType string

const (
	MinifyAll     MinifyType = "all"     // MinifyAll 压缩所有可压缩文件
	MinifyOnlyMin MinifyType = "onlyMin" // MinifyOnlyMin 仅压缩 .min.* 结尾的文件
	MinifyNone    MinifyType = "none"    // MinifyNone 不压缩文件
)

var Config = &config{
	System: system{
		Name:   "Anchor",
		Listen: ":8080",
		Debug:  false,
	},
	Proxy: proxy{
		Github: githubProxy{
			Open:      false,
			Minify:    MinifyOnlyMin,
			Endpoint:  []string{"https://raw.githubusercontent.com/{{user}}/{{repo}}/{{version}}/{{path}}"},
			WhiteList: nil,
		},
		Npm: npmProxy{
			Open:      false,
			Minify:    MinifyOnlyMin,
			Endpoint:  []string{"https://unpkg.com/{{package}}@{{version}}/{{path}}"},
			WhiteList: nil,
		},
		Wp: wordpressProxy{
			PluginOpen:      false,
			ThemeOpen:       false,
			Minify:          MinifyNone,
			PluginWhiteList: nil,
			ThemeWhiteList:  nil,
		},
	},
	CORS: cors{
		AllowOrigins:     []string{"UNSET"},
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "X-Powered-By", "X-Run-By", "X-Timestamp"},
		AllowCredentials: false,
		ExposeHeaders:    nil,
	},
}
