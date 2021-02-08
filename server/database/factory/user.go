package factory

import (
  "fmt"
  "github.com/Pallinder/go-randomdata"
  "github.com/bluele/factory-go/factory"
  "goweibo/app/models"
  "goweibo/core/pkg/numutils"
  "time"
)

var (
  // 头像假数据
  avatars = []string{
    "https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
    "https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
    "https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
    "https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
    "https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
    "https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
  }
)

func userFactory(i int) *factory.Factory {
  n := time.Now()

  u := &models.User{
    Name:            fmt.Sprintf("user-%d", i+1),
    Password:        "123456",
    Email:           randomdata.Email(),
    EmailVerifiedAt: &n,
    Activated:       models.TrueTinyint,
  }

  // 第一个用户是管理员
  if i == 0 {
    u.Name = "admin"
    u.IsAdmin = models.TrueTinyint
    u.Email = "admin@test.com"
  }

  r := numutils.RandInt(0, len(avatars)-1)

  return factory.NewFactory(
    u,
  ).Attr("Avatar", func(args factory.Args) (interface{}, error) {
    return avatars[r], nil
  })
}

func usersTableSeeder() {
  dropAndCreateTable(&models.User{})

  for i := 0; i < 10; i++ {
    u := userFactory(i).MustCreate().(*models.User)
    if err := models.CreateModel(u); err != nil {
      fmt.Printf("mock user error： %v\n", err)
    }
  }
}
