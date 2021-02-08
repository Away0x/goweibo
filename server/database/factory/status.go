package factory

import (
  "github.com/Pallinder/go-randomdata"
  "github.com/bluele/factory-go/factory"
  "goweibo/app/models"
)

var (
  userIDs = []uint{1, 2, 3, 4, 5}
)

func statusFactory(i int) *factory.Factory {
  r := randomdata.Number(0, len(userIDs)-1)
  s := &models.Status{
    UserID: userIDs[r],
  }

  return factory.NewFactory(
    s,
  ).Attr("Content", func(args factory.Args) (interface{}, error) {
    return randomdata.Paragraph(), nil
  })
}

func StatusTableSeeder() {
  dropAndCreateTable(&models.Status{})

  for i := 0; i < 100; i++ {
    status := statusFactory(i).MustCreate().(*models.Status)
    if err := models.CreateModel(status); err != nil {
      panic(err)
    }
  }
}
