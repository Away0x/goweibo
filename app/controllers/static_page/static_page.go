package staticpage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home 主页
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "static_page/home.html", gin.H{})
}

// Help 帮助页
func Help(c *gin.Context) {
	c.HTML(http.StatusOK, "static_page/help.html", gin.H{})
}

// About 关于页
func About(c *gin.Context) {
	c.HTML(http.StatusOK, "static_page/about.html", gin.H{})
}
