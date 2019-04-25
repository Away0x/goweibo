package factory

import (
	"fmt"
	statusModel "gin_weibo/app/models/status"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var (
	userIDs = []uint{1, 2, 3, 4, 5}
)

func statusFactory(i int) *factory.Factory {
	r := randomdata.Number(0, len(userIDs)-1)
	s := &statusModel.Status{
		UserID: userIDs[r],
	}

	return factory.NewFactory(
		s,
	).Attr("Content", func(args factory.Args) (interface{}, error) {
		return randomdata.Paragraph(), nil
	})
}

// StatusTableSeeder -
func StatusTableSeeder(needCleanTable bool) {
	if needCleanTable {
		DropAndCreateTable(&statusModel.Status{})
	}

	for i := 0; i < 100; i++ {
		status := statusFactory(i).MustCreate().(*statusModel.Status)
		if err := status.Create(); err != nil {
			fmt.Printf("mock status error： %v\n", err)
		}
	}
}
