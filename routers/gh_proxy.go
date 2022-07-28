package routers

import (
	"github.com/AH-dark/Anchor/controllers"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/gin-gonic/gin"
)

// InitGithubProxy 初始化 GitHub 代理功能
func InitGithubProxy(r *gin.RouterGroup) {
	utils.Log().Info("GitHub 代理已开启")

	{
		r.GET(":user/:repo/:version/*path", controllers.GithubRawFileProxy)
	}
}
