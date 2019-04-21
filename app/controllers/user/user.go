package user

import (
	"github.com/gin-gonic/gin"

	"gin_weibo/app/auth"
	"gin_weibo/app/models"

	"gin_weibo/app/controllers"
	"gin_weibo/app/policies"
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
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	user, err := auth.GetUserFromContextOrDataBase(c, id)
	if err != nil || user == nil {
		controllers.Render404(c)
		return
	}

	controllers.Render(c, "user/show.html", gin.H{
		"userData": viewmodels.NewUserViewModelSerializer(user),
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
func Edit(c *gin.Context, currentUser *models.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 只能查看自己的编辑页面
	if ok := policies.UserPlolicyUpdate(c, currentUser, id); !ok {
		return
	}

	controllers.Render(c, "user/edit.html", gin.H{
		"userData": viewmodels.NewUserViewModelSerializer(currentUser),
	})
}

// Update 编辑用户
func Update(c *gin.Context, currentUser *models.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 只能更新自己
	if ok := policies.UserPlolicyUpdate(c, currentUser, id); !ok {
		return
	}

	// 验证参数和更新用户
	userUpdateForm := &userRequest.UserUpdateForm{
		Name:                 c.PostForm("name"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	errors := userUpdateForm.ValidateAndSave(currentUser)

	if len(errors) != 0 {
		flash.SaveValidateMessage(c, errors)
		controllers.RedirectToUserEditPage(c, currentUser.GetIDstring())
		return
	}

	flash.NewSuccessFlash(c, "个人资料更新成功！")
	controllers.RedirectToUserShowPage(c, currentUser)
}

// Destroy 删除用户
func Destroy(c *gin.Context) {

}
