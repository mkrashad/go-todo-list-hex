package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/api-gw/ctxutils"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(ctxutils.SetLogger(c.Request.Context()))
		c.Next()
	}
}
