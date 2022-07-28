package routers

import (
	"github.com/AH-dark/Anchor/controllers"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/gin-gonic/gin"
)

// InitNpmProxy 初始化 Npm 代理功能
func InitNpmProxy(r *gin.RouterGroup) {
	utils.Log().Info("NPM 代理已开启")

	{
		r.GET("@:user/:package", controllers.NpmRedirect)
		r.GET(":package", controllers.NpmRedirect)

		r.GET(":package/*path", controllers.NpmRawFileProxy)
		r.GET("@:user/:package/*path", controllers.NpmRawFileProxy)
	}
}
