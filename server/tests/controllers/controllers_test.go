package controllers_test

import (
  "github.com/gavv/httpexpect/v2"
  "goweibo/bootstrap"
  "goweibo/core"
  "net/http"
  "os"
  "testing"
)

func TestMain(m *testing.M) {
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
  bootstrap.SetupServer()
  return
}

func after() {}

func apiClient(t *testing.T) *httpexpect.Expect {
  engine := core.GetApplication().Echo

  return httpexpect.WithConfig(httpexpect.Config{
    BaseURL: "/api",
    Client: &http.Client{
      Transport: httpexpect.NewBinder(engine),
      Jar:       httpexpect.NewJar(),
    },
    Reporter: httpexpect.NewAssertReporter(t),
    Printers: []httpexpect.Printer{
      // log
      // httpexpect.NewDebugPrinter(t, true),
    },
  })
}

func getOKApiJSon(r *httpexpect.Request) *httpexpect.Object {
  return r.Expect().Status(http.StatusOK).JSON().Object()
}
