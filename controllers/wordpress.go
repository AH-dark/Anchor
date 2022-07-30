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

func WordpressRawFileProxy(proxyType services.WpProxyType) gin.HandlerFunc {
	var endpoint = ""
	switch proxyType {
	case services.WpProxyTypeTheme:
		endpoint = "https://themes.svn.wordpress.org/{{name}}/{{version}}/"
	case services.WpProxyTypePlugin:
		endpoint = "https://plugins.svn.wordpress.org/{{name}}/tags/{{version}}/"
	}

	return func(c *gin.Context) {
		if endpoint == "" {
			c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		// Parse param data
		name := c.Param("name")
		version := c.Param("version")
		path := c.Param("path")

		// check in whitelist
		if !services.CheckWordpressWhitelist(proxyType, name) {
			c.String(http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}

		// get body data
		data := services.GetWordpressRawFile(endpoint, name, version, path)
		if data == nil {
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		// compress
		contentType := mime.TypeByExtension("." + utils.Extension(path))
		if conf.Config.Proxy.Wp.Minify != conf.MinifyNone && compress.CanBeCompressed(path, contentType, conf.MinifyAll) {
			data = compress.Bytes(data, contentType)
			c.Header("X-Anchor-Minify", "true")
		} else {
			c.Header("X-Anchor-Minify", "false")
		}

		c.Data(http.StatusOK, contentType, data)
		c.Next()
	}
}
