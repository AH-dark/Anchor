package routers

import (
	"github.com/AH-dark/Anchor/controllers"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/gin-gonic/gin"
)

func InitGithubProxy(r *gin.RouterGroup) {
	utils.Log().Info("GitHub 代理已开启")

	{
		r.Any(":user/:repo/:version/*path", controllers.GithubRawFileProxy)
	}
}
