package staticpage

import (
	"gin_weibo/app/auth"
	"gin_weibo/app/controllers"
	statusModel "gin_weibo/app/models/status"
	viewmodels "gin_weibo/app/view_models"
	"gin_weibo/pkg/pagination"
	"gin_weibo/routes/named"

	"github.com/gin-gonic/gin"
)

// Home 主页
func Home(c *gin.Context) {
	currentUser, err := auth.GetCurrentUserFromContext(c)
	if err != nil {
		controllers.Render(c, "static_page/home.html", gin.H{})
		return
	}

	// 获取分页参数
	statusesAllLength, _ := statusModel.GetUserAllStatusCount(int(currentUser.ID))
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, 10, statusesAllLength)
	if currentPage > pageTotalCount {
		controllers.Redirect(c, named.G("root")+"?page=1", false)
		return
	}
	// 获取用户的微博
	statuses, _ := statusModel.GetUserStatus(int(currentUser.ID), offset, limit)
	statusesViewModels := make([]*viewmodels.StatusViewModel, 0)
	for _, s := range statuses {
		statusesViewModels = append(statusesViewModels, viewmodels.NewStatusViewModelSerializer(s))
	}

	controllers.Render(c, "static_page/home.html",
		pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount, gin.H{
			"statuses":       statusesViewModels,
			"statusesLength": statusesAllLength,
		}))
}

// Help 帮助页
func Help(c *gin.Context) {
	controllers.Render(c, "static_page/help.html", gin.H{})
}

// About 关于页
func About(c *gin.Context) {
	controllers.Render(c, "static_page/about.html", gin.H{})
}
