package conf

var Config = &config{
	System: system{
		Name:  "Anchor",
		Port:  ":8080",
		Debug: false,
	},
	Proxy: proxy{
		Github: githubProxy{
			Open:      false,
			Minify:    false,
			Endpoint:  []string{"https://raw.githubusercontent.com/{{user}}/{{repo}}/{{version}}/{{path}}"},
			WhiteList: nil,
		},
		Npm: npmProxy{
			Open:      false,
			Minify:    false,
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
