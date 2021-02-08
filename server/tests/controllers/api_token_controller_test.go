package controllers_test

import (
  "github.com/gavv/httpexpect/v2"
  "goweibo/app/requests"
  "goweibo/core/constants"
  "goweibo/tests"
  "testing"
)

func TestApiTokenControllerStore(t *testing.T) {
  var (
    client = apiClient(t)
    url = "/token/store"
    resp *httpexpect.Object
  )
  u := tests.CreateUserModel(t)

  // fail
  resp = getOKApiJSon(client.POST(url))
  resp.Value("code").Equal(constants.RequestErrorCode)

  // ok
  resp = getOKApiJSon(client.POST(url).WithJSON(requests.UserLogin{
    Email: u.Email,
    Password: tests.DefaultUserPassword,
  }))
  resp.Value("code").Equal(constants.SuccessCode)
}
