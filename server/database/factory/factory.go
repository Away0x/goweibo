package factory

import (
  "fmt"
  "goweibo/core"
)

func dropAndCreateTable(table interface{}) {
	core.GetDefaultConnectionEngine().Migrator().DropTable(table)
	core.GetDefaultConnectionEngine().Migrator().CreateTable(table)
}

// Run run database factory
func Run() {
  usersTableSeeder()
	fmt.Println("database.factory runing")
}
