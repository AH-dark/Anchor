package routers

import (
	"github.com/AH-dark/Anchor/controllers"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/gin-gonic/gin"
)

func InitGithubProxy(r *gin.Engine) {
	utils.Log().Info("GitHub 代理已开启")

	g := r.Group("/gh/")
	{
		g.Any(":user/:repo/:version/*path", controllers.GithubRawFileProxy)
	}
}
