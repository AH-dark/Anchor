package middleware

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS 初始化跨域配置
func CORS() gin.HandlerFunc {
	if conf.Config.CORS.AllowOrigins[0] != "UNSET" {
		return cors.New(cors.Config{
			AllowOrigins:     conf.Config.CORS.AllowOrigins,
			AllowMethods:     conf.Config.CORS.AllowMethods,
			AllowHeaders:     conf.Config.CORS.AllowHeaders,
			AllowCredentials: conf.Config.CORS.AllowCredentials,
			ExposeHeaders:    conf.Config.CORS.ExposeHeaders,
		})
	}

	return cors.Default()
}
