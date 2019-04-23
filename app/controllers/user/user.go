package user

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"gin_weibo/app/auth"
	"gin_weibo/app/models"

	"gin_weibo/app/controllers"
	"gin_weibo/app/helpers"
	"gin_weibo/app/policies"
	userRequest "gin_weibo/app/requests/user"
	"gin_weibo/app/services"
	viewmodels "gin_weibo/app/view_models"
	"gin_weibo/pkg/flash"
	"gin_weibo/pkg/pagination"
)

// Index 用户列表
func Index(c *gin.Context, currentUser *models.User) {
	var (
		m               = models.User{}
		defaultPageLine = 10
	)

	allUserCount := m.AllCount()
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, defaultPageLine, allUserCount)

	if currentPage > pageTotalCount {
		// controllers.Render404(c)
		controllers.RedirectToUserIndexPage(c, "1") // 超出就跳转到第一页
		return
	}

	users := services.UserListService(offset, limit)

	controllers.Render(c, "user/index.html",
		pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount, gin.H{
			"users": users,
		}))
}

// Create 创建用户页面
func Create(c *gin.Context) {
	controllers.Render(c, "user/create.html", gin.H{})
}

// Show 用户详情
func Show(c *gin.Context, currentUser *models.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 如果要看的就是当前用户，那么就不用再去数据库中获取了
	user := currentUser
	if id != int(currentUser.ID) {
		user = &models.User{}
		err = user.Get(id)
	}

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

	// auth.Login(c, user)
	// flash.NewSuccessFlash(c, "欢迎，您将在这里开启一段新的旅程~")
	// controllers.RedirectToUserShowPage(c, user)

	sendConfirmEmail(user)
	flash.NewSuccessFlash(c, "验证邮件已发送到你的注册邮箱上，请注意查收。")
	controllers.RedirectToRootPage(c)
}

// Edit 编辑用户页面
func Edit(c *gin.Context, currentUser *models.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 只能查看自己的编辑页面
	if ok := policies.UserPolicyUpdate(c, currentUser, id); !ok {
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
	if ok := policies.UserPolicyUpdate(c, currentUser, id); !ok {
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

// Destory 删除用户
func Destory(c *gin.Context, currentUser *models.User) {
	page := c.DefaultQuery("page", "1")

	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 是否有删除权限
	if ok := policies.UserPolicyDestory(c, currentUser, id); !ok {
		return
	}

	// 删除用户
	user := &models.User{}
	if err = user.Delete(id); err != nil {
		flash.NewDangerFlash(c, "删除失败: "+err.Error())
		controllers.RedirectToUserIndexPage(c, page)
		return
	}

	flash.NewSuccessFlash(c, "成功删除用户！")
	controllers.RedirectToUserIndexPage(c, page)
}

// ConfirmEmail : 邮箱验证
func ConfirmEmail(c *gin.Context) {
	token := c.Param("token")

	user := &models.User{}
	err := user.GetByActivationToken(token)
	if user == nil || err != nil {
		controllers.Render404(c)
		return
	}

	// 更新用户
	user.Activated = models.TrueTinyint
	user.ActivationToken = ""
	user.EmailVerifiedAt = time.Now()
	if err = user.Update(false); err != nil {
		flash.NewSuccessFlash(c, "用户激活失败: "+err.Error())
		controllers.RedirectToRootPage(c)
		return
	}

	auth.Login(c, user)
	flash.NewSuccessFlash(c, "恭喜你，激活成功！")
	controllers.RedirectToUserShowPage(c, user)
}

// -------------- private
func sendConfirmEmail(u *models.User) {
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/confirm.html"
	confirmURL := controllers.CreateSignupConfirmURL(u)

	if err := helpers.SendMail([]string{u.Email}, subject, tpl, gin.H{"confirmURL": confirmURL}); err != nil {
		fmt.Printf("send mail %v\n", err)
	}
}
