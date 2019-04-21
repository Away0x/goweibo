package middleware

import (
	"gin_weibo/app/auth"

	"github.com/gin-gonic/gin"
)

// Auth : auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth.SaveUserToContext(c)

		c.Next()
	}
}
