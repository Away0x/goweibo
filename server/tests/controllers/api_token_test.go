package controllers_test

import (
  "github.com/gavv/httpexpect/v2"
  "goweibo/app/requests"
  "goweibo/core/constants"
  "testing"
)

func TestApiTokenStore(t *testing.T) {
  var (
    client = apiClient(t)
    url = "/token/store"
    resp *httpexpect.Object
  )

  // fail
  resp = getOKApiJSon(client.POST(url))
  resp.Value("code").Equal(constants.RequestErrorCode)

  // ok
  resp = getOKApiJSon(client.POST(url).WithJSON(requests.UserLogin{
    Email: "1@qq.com",
    Password: "123456",
  }))
  resp.Value("code").Equal(constants.SuccessCode)
}
