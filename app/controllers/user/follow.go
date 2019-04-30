package user

import (
	"gin_weibo/app/controllers"
	followerModel "gin_weibo/app/models/follower"
	userModel "gin_weibo/app/models/user"
	viewmodels "gin_weibo/app/view_models"
	"gin_weibo/pkg/pagination"
	"gin_weibo/routes/named"

	"github.com/gin-gonic/gin"
)

// Followings 用户关注者列表
func Followings(c *gin.Context, currentUser *userModel.User) {
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
	followingsLength, _ := followerModel.FollowingsCount(id)
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, 10, followingsLength)
	if currentPage > pageTotalCount {
		controllers.Redirect(c, named.G("users.followings")+"?page=1", false)
		return
	}

	// 获取关注者
	followings, _ := followerModel.Followings(id, offset, limit)
	usersViewModels := make([]*viewmodels.UserViewModel, 0)
	for _, u := range followings {
		usersViewModels = append(usersViewModels, viewmodels.NewUserViewModelSerializer(u))
	}

	controllers.Render(c, "user/show_follow.html",
		pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount, gin.H{
			"title":    user.Name + " 关注的人",
			"userData": viewmodels.NewUserViewModelSerializer(user),
			"users":    usersViewModels,
		}))
}

// Followers 用户粉丝列表
func Followers(c *gin.Context, currentUser *userModel.User) {
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
	followersLength, _ := followerModel.FollowersCount(id)
	offset, limit, currentPage, pageTotalCount := controllers.GetPageQuery(c, 10, followersLength)
	if currentPage > pageTotalCount {
		controllers.Redirect(c, named.G("users.followers")+"?page=1", false)
		return
	}

	// 获取关注者
	followers, _ := followerModel.Followers(id, offset, limit)
	usersViewModels := make([]*viewmodels.UserViewModel, 0)
	for _, u := range followers {
		usersViewModels = append(usersViewModels, viewmodels.NewUserViewModelSerializer(u))
	}

	controllers.Render(c, "user/show_follow.html",
		pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount, gin.H{
			"title":    user.Name + " 的粉丝",
			"userData": viewmodels.NewUserViewModelSerializer(user),
			"users":    usersViewModels,
		}))
}
