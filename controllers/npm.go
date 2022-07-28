package controllers

import (
	"github.com/AH-dark/Anchor/pkg/compress"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/AH-dark/Anchor/services"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"strings"
)

func parseNpmPackageInfo(packageInfo string) (string, string) {
	i := strings.Split(packageInfo, "@")

	switch len(i) {
	default:
		return "", ""
	case 1:
		return i[0], ""
	case 2:
		return i[0], i[1]
	}
}

func NpmRawFileProxy(c *gin.Context) {
	user := c.Param("user")
	pkg, version := parseNpmPackageInfo(c.Param("package"))
	path := c.Param("path")

	// 校验
	if !services.CheckNpmWhiteList(user, pkg) {
		c.String(http.StatusForbidden, http.StatusText(http.StatusForbidden))
		return
	}

	contentType := mime.TypeByExtension("." + utils.Extension(path))
	compressMode := compress.CanBeCompressed(path, contentType, conf.Config.Proxy.Npm.Minify)

	// 获取文件
	data := services.GetNpmRawFile(user, pkg, version, path)
	if data == nil {
		if compressMode {
			path = utils.RemoveMinSuffix(path)
			utils.Log().Debug("原路径请求失败，即将使用新路径重新请求：%s", path)
			data = services.GetNpmRawFile(user, pkg, version, path)
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
