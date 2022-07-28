package controllers

import (
	"github.com/AH-dark/Anchor/pkg/compress"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/AH-dark/Anchor/services"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
)

func GithubRawFileProxy(c *gin.Context) {
	user := c.Param("user")
	repo := c.Param("repo")
	version := c.Param("version")
	path := c.Param("path")

	// 校验
	if !services.CheckGithubWhiteList(user, repo) {
		c.String(http.StatusForbidden, http.StatusText(http.StatusForbidden))
		return
	}

	contentType := mime.TypeByExtension("." + utils.Extension(path))
	compressMode := compress.CanBeCompressed(path, contentType, conf.Config.Proxy.Github.Minify)

	// 获取文件
	data := services.GetGithubRawFile(user, repo, version, path)
	if data == nil {
		if compressMode {
			path = utils.RemoveMinSuffix(path)
			utils.Log().Debug("原路径请求失败，即将使用新路径重新请求：%s", path)
			data = services.GetGithubRawFile(user, repo, version, path)
			if data == nil {
				c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
				return
			}
		} else {
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
	}

	// 压缩文件
	if compressMode {
		data = compress.Bytes(data, contentType)
		c.Header("X-Anchor-Minify", "true")
	} else {
		c.Header("X-Anchor-Minify", "false")
	}

	c.Data(http.StatusOK, contentType, data)
	c.Next()
}
