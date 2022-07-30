package routers

import (
	"github.com/AH-dark/Anchor/controllers"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/AH-dark/Anchor/services"
	"github.com/gin-gonic/gin"
)

// InitWordpressProxy 初始化 WordPress 代理功能
func InitWordpressProxy(r *gin.RouterGroup) {
	utils.Log().Info("WordPress 代理已开启")

	{
		if conf.Config.Proxy.Wp.ThemeOpen {
			r.GET("theme/:name/:version/*path", controllers.WordpressRawFileProxy(services.WpProxyTypeTheme))
		}
		if conf.Config.Proxy.Wp.PluginOpen {
			r.GET("plugin/:name/:version/*path", controllers.WordpressRawFileProxy(services.WpProxyTypePlugin))
		}
	}
}
