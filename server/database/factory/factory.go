package factory

import (
	"fmt"
	"goweibo/bootstrap"
	"goweibo/core"
)

func dropAndCreateTable(table interface{}) {
	core.GetDefaultConnectionEngine().Migrator().DropTable(table)
	core.GetDefaultConnectionEngine().Migrator().CreateTable(table)
}

// Run run database factory
func Run() {
	bootstrap.SetupDB()

  usersTableSeeder()
	fmt.Println("database.factory runing")
}
