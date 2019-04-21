package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin_weibo/app/auth"
	"gin_weibo/app/models"

	"gin_weibo/app/controllers"
	userRequest "gin_weibo/app/requests/user"
	viewmodels "gin_weibo/app/view_models"
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

	user := &models.User{}
	err = user.Get(id)
	if err != nil {
		c.String(http.StatusOK, "用户获取错误 %v", err)
		return
	}

	controllers.Render(c, "user/show.html", gin.H{
		"userData": viewmodels.NewUserViewModelSerializer(user, 140),
	})
}

// Store 保存用户
func Store(c *gin.Context) {
	// 验证参数和创建用户
	userCreateForm := &userRequest.UserCreateForm{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	user, errors := userCreateForm.ValidateAndSave()

	if len(errors) != 0 || user == nil {
		flash.SaveValidateMessage(c, errors)
		controllers.RedirectToUserCreatePage(c)
		return
	}

	auth.Login(c, user)
	flash.NewSuccessFlash(c, "欢迎，您将在这里开启一段新的旅程~")
	controllers.RedirectToUserShowPage(c, user)
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
