package controllers

import (
	"github.com/AH-dark/Anchor/pkg/compress"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/AH-dark/Anchor/services"
	"github.com/gin-gonic/gin"
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

	// 获取文件
	data, contentType := services.GetGithubRawFile(user, repo, version, path)
	if data == nil {
		c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	// 压缩文件
	switch conf.Config.Proxy.Github.Minify {
	case conf.MinifyAll:
		data = compress.CompressBytes(data, contentType)
		c.Header("X-Anchor-Minify", "true")
	case conf.MinifyOnlyMin:
		if utils.FileHasMinSuffix(path) {
			data = compress.CompressBytes(data, contentType)
			c.Header("X-Anchor-Minify", "true")
		}
	default:
		c.Header("X-Anchor-Minify", "false")
	}

	c.Data(http.StatusOK, contentType, data)
	c.Next()
}
