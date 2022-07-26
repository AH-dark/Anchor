package conf

const (
	MinifyAll     = "all"     // MinifyAll 压缩所有可压缩文件
	MinifyOnlyMin = "onlyMin" // MinifyOnlyMin 仅压缩 .min.* 结尾的文件
	MinifyNone    = "none"    // MinifyNone 不压缩文件
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
	},
	CORS: cors{
		AllowOrigins:     []string{"UNSET"},
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Content-Type"},
		AllowCredentials: false,
		ExposeHeaders:    nil,
	},
}
