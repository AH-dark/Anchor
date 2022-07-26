package bootstrap

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/gin-gonic/gin"
)

func Init(path string) {
	conf.Init(path)

	// Debug 关闭时，切换为生产模式
	if !conf.Config.System.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
