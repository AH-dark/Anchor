package controllers

import (
	"github.com/gin-gonic/gin"
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
	packageName, version := parseNpmPackageInfo(c.Param("package"))
	path := c.Param("path")

	c.String(http.StatusOK, "Hi!\nuser: %s\npackage: %s\nversion: %s\npath: %s", user, packageName, version, path)
}
