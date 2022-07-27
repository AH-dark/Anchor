package middleware

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CustomHeaders() gin.HandlerFunc {
	name := conf.Config.System.Name

	return func(c *gin.Context) {
		c.Header("X-Powered-By", "AH-dark/Anchor")
		c.Header("X-Run-By", name)
		c.Header("X-Timestamp", strconv.FormatInt(time.Now().UnixMicro(), 10))
	}
}
