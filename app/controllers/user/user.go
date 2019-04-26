package user

import (
	"time"

	"github.com/gin-gonic/gin"

	"gin_weibo/app/auth"
	"gin_weibo/app/models"
	statusModel "gin_weibo/app/models/status"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/routes/named"

	"gin_weibo/app/controllers"
	"gin_weibo/app/policies"
	userRequest "gin_weibo/app/requests/user"
	"gin_weibo/app/services"
	viewmodels "gin_weibo/app/view_models"
	"gin_weibo/pkg/flash"
	"gin_weibo/pkg/pagination"
)

// Index 用户列表
func Index(c *gin.Context, currentUser *userModel.User) {
	defaultPageLine := 10

	allUserCount, err := userModel.AllCount()
	if err != nil {
		flash.NewDangerFlash(c, "获取用户数据失败: "+err.Error())
		controllers.Redirect(c, named.G("users.index")+"?page=1", false)
		return
	}
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, defaultPageLine, allUserCount)

	if currentPage > pageTotalCount {
		controllers.Redirect(c, named.G("users.index")+"?page=1", false)
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
func Show(c *gin.Context, currentUser *userModel.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 如果要看的就是当前用户，那么就不用再去数据库中获取了
	user := currentUser
	if id != int(currentUser.ID) {
		user, err = userModel.Get(id)
	}

	if err != nil || user == nil {
		controllers.Render404(c)
		return
	}

	// 获取分页参数
	statusesAllLength, _ := statusModel.GetUserAllStatusCount(int(user.ID))
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, 10, statusesAllLength)
	if currentPage > pageTotalCount {
		controllers.Redirect(c, named.G("users.show")+"?page=1", false)
		return
	}

	// 获取用户的微博
	statuses, _ := statusModel.GetUserStatus(int(user.ID), offset, limit)
	statusesViewModels := make([]*viewmodels.StatusViewModel, 0)
	for _, s := range statuses {
		statusesViewModels = append(statusesViewModels, viewmodels.NewStatusViewModelSerializer(s))
	}

	controllers.Render(c, "user/show.html",
		pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount, gin.H{
			"userData":         viewmodels.NewUserViewModelSerializer(user),
			"statuses":         statusesViewModels,
			"statusesLength":   statusesAllLength,
			"followingsLength": 0,
			"followersLength":  0,
		}))
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
		controllers.RedirectRouter(c, "users.create")
		return
	}

	if err := sendConfirmEmail(user); err != nil {
		flash.NewDangerFlash(c, "验证邮件发送失败: "+err.Error())
	} else {
		flash.NewSuccessFlash(c, "验证邮件已发送到你的注册邮箱上，请注意查收。")
	}

	controllers.RedirectRouter(c, "root")
}

// Edit 编辑用户页面
func Edit(c *gin.Context, currentUser *userModel.User) {
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
func Update(c *gin.Context, currentUser *userModel.User) {
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
		controllers.RedirectRouter(c, "users.edit", currentUser.ID)
		return
	}

	flash.NewSuccessFlash(c, "个人资料更新成功！")
	controllers.RedirectRouter(c, "users.show", currentUser.ID)
}

// Destory 删除用户
func Destory(c *gin.Context, currentUser *userModel.User) {
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
	if err = userModel.Delete(id); err != nil {
		flash.NewDangerFlash(c, "删除失败: "+err.Error())
	} else {
		flash.NewSuccessFlash(c, "成功删除用户！")
	}

	controllers.Redirect(c, named.G("users.index")+"?page="+page, false)
}

// ConfirmEmail : 邮箱验证
func ConfirmEmail(c *gin.Context) {
	token := c.Param("token")

	user, err := userModel.GetByActivationToken(token)
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
		controllers.RedirectRouter(c, "root")
		return
	}

	auth.Login(c, user)
	flash.NewSuccessFlash(c, "恭喜你，激活成功！")
	controllers.RedirectRouter(c, "users.show", user.ID)
}
