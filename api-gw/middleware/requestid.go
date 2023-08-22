package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mkrashad/go-todo/api-gw/ctxutils"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.NewString()
		c.Request = c.Request.WithContext(ctxutils.SetRequestId(c.Request.Context(), requestId))
		c.Next()
	}
}

