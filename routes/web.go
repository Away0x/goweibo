package routes

import (
	"github.com/gin-gonic/gin"

	staticpage "gin_weibo/app/controllers/static_page"
	"gin_weibo/app/controllers/user"
)

func registerWeb(g *gin.Engine) {
	// static page
	g.GET("/", staticpage.Home)
	g.GET("/help", staticpage.Help)
	g.GET("/about", staticpage.About)

	// user
	g.GET("/signup", user.Create)
}
