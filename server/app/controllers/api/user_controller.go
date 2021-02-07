package api

import (
  "goweibo/app/auth"
  "goweibo/app/models"
  "goweibo/app/requests"
  "goweibo/app/serializers"
  "goweibo/app/services"
  "goweibo/core/context"
  "goweibo/core/errno"
)

type IUserController interface {
  // Index 获取用户列表
  Index(*context.AppContext, *auth.TokenAuthInfo) error
  // Show 获取用户详情
  Show(*context.AppContext, *auth.TokenAuthInfo, *models.User) error
  // 创建用户
  Create(*context.AppContext) error
}

type UserController struct {
  UserServices services.IUserServices
}

func NewUserController(s services.IUserServices) IUserController {
  return &UserController{UserServices: s}
}

// Index 获取用户列表
// @Summary 获取用户列表
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} context.CommonResponse{data=[]serializers.UserList}
// @Security ApiKeyAuth
// @Router /user [get]
func (u *UserController) Index(c *context.AppContext, a *auth.TokenAuthInfo) error {
  users, err := u.UserServices.List()
  if err != nil {
    return errno.DatabaseErr.WithErr(err)
  }
  return c.AWSuccessJSON(serializers.NewUserListSerializer(users))
}

// Show 获取用户信息
// @Summary 获取用户信息
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "用户 id"
// @Success 200 {object} context.CommonResponse{data=serializers.User}
// @Security ApiKeyAuth
// @Router /user/{id} [get]
func (*UserController) Show(c *context.AppContext, a *auth.TokenAuthInfo, u *models.User) error {
  return c.AWSuccessJSON(serializers.NewUserSerializer(u))
}

// Create 创建用户
// @Summary 创建用户
// @Tags User
// @Accept json
// @Produce json
// @Param json body requests.CreateUser true "用户信息"
// @Success 200 {object} context.CommonResponse{data=serializers.User}
// @Security ApiKeyAuth
// @Router /user [post]
func (*UserController) Create(c *context.AppContext) (err error) {
  req := new(requests.CreateUser)
  if err = c.AWBindValidatorStruct(req); err != nil {
    return err
  }

  u, err := req.Create()
  if err != nil {
    return err
  }

  return c.AWSuccessJSON(serializers.NewUserSerializer(u))
}

