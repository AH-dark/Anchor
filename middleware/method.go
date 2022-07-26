package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var methodAllow = []string{
	"GET",
	"HEAD",
	"OPTIONS",
}

func MethodControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range methodAllow {
			if c.Request.Method == v {
				c.Next()
				return
			}
		}

		c.String(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}
