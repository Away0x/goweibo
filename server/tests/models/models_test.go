package models_test

import (
  "goweibo/app/models"
  "goweibo/bootstrap"
  "goweibo/core"
  "os"
  "testing"
)

func TestMain(m *testing.M)  {
  if err := before(); err != nil {
    panic(err)
  }
  m.Run()
  after()
}

func before() (err error) {
  os.Chdir("../..")

  bootstrap.SetupConfig("config/test.yaml", "yaml")
  bootstrap.SetupDB()

  err = clearDatabase(&models.User{})
  return
}

func after() {}

func clearDatabase(tables ...interface{}) (err error) {
  for _, t := range tables {
    err = core.GetDefaultConnectionEngine().Migrator().DropTable(t)
    err = core.GetDefaultConnectionEngine().Migrator().CreateTable(t)
  }

  return
}
