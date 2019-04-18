package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create 创建用户
func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "user/create.html", gin.H{})
}
