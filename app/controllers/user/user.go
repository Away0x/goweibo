package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin_weibo/app/models"
	viewmodels "gin_weibo/app/view_models"

	"gin_weibo/app/controllers"
	"gin_weibo/pkg/flash"
)

// Index 用户列表
func Index(c *gin.Context) {
	controllers.Render(c, "user/index.html", gin.H{
		"my": "user index",
	})
}

// Create 创建用户页面
func Create(c *gin.Context) {
	controllers.Render(c, "user/create.html", gin.H{})
}

// Show 用户详情
func Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusOK, "参数错误 %v", err)
		return
	}

	m := &models.User{}
	user, err := m.Get(id)
	if err != nil {
		c.String(http.StatusOK, "用户获取错误 %v", err)
		return
	}

	c.HTML(http.StatusOK, "user/show.html", viewmodels.NewUserViewModelSerializer(user, 140))
}

// Store 保存用户
func Store(c *gin.Context) {
	flash.NewSuccessFlash(c, "啦啦啦啦写入 flash 成功啦")

	// c.Request.Method = "POST"
	// c.Request.URL.Path = "/users"
	// r.HandleContext(c)
	controllers.Redirect(c, "http://localhost:8888/users/create")
}

// Edit 编辑用户页面
func Edit(c *gin.Context) {

}

// Update 编辑用户
func Update(c *gin.Context) {

}

// Destroy 删除用户
func Destroy(c *gin.Context) {

}
