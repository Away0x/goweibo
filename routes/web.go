package routes

import (
	"github.com/gin-gonic/gin"

	staticpage "gin_weibo/app/controllers/static_page"
)

func registerWeb(g *gin.Engine) {
	// root
	g.GET("/", staticpage.Home)
	g.GET("/help", staticpage.Help)
	g.GET("/about", staticpage.About)
}
