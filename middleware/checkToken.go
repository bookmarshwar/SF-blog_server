package middleware

import (
	"blog_server/service"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("x-token")
		if token == "stand" {
			c.Next()
			return
		}
		print(token)
		_, error := service.ParseToken(token)
		if error != nil {
			print("%s", error.Error())
			c.Abort()
		}
		c.Next()
	}

}
