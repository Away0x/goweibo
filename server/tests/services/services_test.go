package services

import (
  "goweibo/bootstrap"
  "goweibo/tests"
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
  err = os.Chdir("../..")
  bootstrap.SetupConfig("config/test.yaml", "yaml")
  bootstrap.SetupDB()
  err = tests.ResetDatabase()
  return
}

func after() {}
