package factory

import (
	"fmt"
	"gin_weibo/app/models"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/pkg/utils"

	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var (
	// 头像假数据
	avatars = []string{
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/s5ehp11z6s.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/Lhd1SHqu86.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/LOnMrqbHJn.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/xAuDMxteQy.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/ZqM7iaP4CR.png",
		"https://iocaffcdn.phphub.org/uploads/images/201710/14/1/NDnzMutoxX.png",
	}
)

func userFactory(i int) *factory.Factory {
	u := &userModel.User{
		Password:        "123456",
		EmailVerifiedAt: time.Now(),
		Activated:       models.TrueTinyint,
		RememberToken:   string(utils.RandomCreateBytes(10)),
	}
	// 第一个用户是管理员
	if i == 0 {
		u.IsAdmin = models.TrueTinyint
	}

	r := utils.RandInt(0, len(avatars)-1)

	return factory.NewFactory(
		u,
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return fmt.Sprintf("user-%d", i+1), nil
	}).Attr("Avatar", func(args factory.Args) (interface{}, error) {
		return avatars[r], nil
	}).Attr("Email", func(args factory.Args) (interface{}, error) {
		if i == 0 {
			return "1@test.com", nil
		}
		return randomdata.Email(), nil
	})
}

// UsersTableSeeder -
func UsersTableSeeder(needCleanTable bool) {
	if needCleanTable {
		DropAndCreateTable(&userModel.User{})
	}

	for i := 0; i < 100; i++ {
		user := userFactory(i).MustCreate().(*userModel.User)
		if err := user.Create(); err != nil {
			fmt.Printf("mock user error： %v\n", err)
		}
	}
}
