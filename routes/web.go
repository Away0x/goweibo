package routes

import (
	"github.com/gin-gonic/gin"

	"gin_weibo/app/controllers/home"
)

func registerWeb(g *gin.Engine) {
	// root
	g.GET("/", home.Index)
	g.GET("/2", home.Index2)
}
