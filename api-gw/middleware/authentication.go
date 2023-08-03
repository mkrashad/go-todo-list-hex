package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/api-gw/handler/auth"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			token := c.Request.Header.Get("Authorization")
			if token == "" {
					log.Println("Authorization header is empty")
					c.AbortWithStatus(401)
					return
			}

			claims, err := auth.ParseToken(token)
			if err != nil {
					log.Printf("Parsing token error: %v\n", err.Error())
					c.AbortWithStatus(401)
					return
			}

			c.Set("userId", claims["uid"])
			c.Next()
	}
}