package factory

import (
  "fmt"
  "github.com/Pallinder/go-randomdata"
  "github.com/bluele/factory-go/factory"
  "goweibo/app/models"
)

func userFactory(i int) *factory.Factory {
  u := &models.User{
    Name: fmt.Sprintf("user-%d", i + 1),
    Password: "123456",
    Email: randomdata.Email(),
  }

  return factory.NewFactory(u)
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
