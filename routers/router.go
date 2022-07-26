package routers

import (
	"github.com/AH-dark/Anchor/middleware"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())
	r.Use(middleware.Gzip())
	r.Use(middleware.MethodControl())

	// GitHub proxy
	if conf.Config.Proxy.Github.Open {
		InitGithubProxy(r)
	}

	// Npm proxy
	if conf.Config.Proxy.Npm.Open {
		InitNpmProxy(r)
	}

	return r
}
