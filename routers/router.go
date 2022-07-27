package routers

import (
	"github.com/AH-dark/Anchor/middleware"
	"github.com/AH-dark/Anchor/pkg/conf"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"time"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())
	r.Use(middleware.Gzip())
	r.Use(middleware.MethodControl())
	r.Use(middleware.CustomHeaders())

	store := persist.NewMemoryStore(1 * time.Minute)

	// GitHub proxy
	if conf.Config.Proxy.Github.Open {
		gh := r.Group("/gh/")
		gh.Use(cache.CacheByRequestURI(store, 30*24*time.Hour))

		InitGithubProxy(gh)
	}

	// Npm proxy
	if conf.Config.Proxy.Npm.Open {
		npm := r.Group("/npm/")
		npm.Use(cache.CacheByRequestURI(store, 3*24*time.Hour))

		InitNpmProxy(npm)
	}

	return r
}
