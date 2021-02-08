package api

import (
  "fmt"
  "goweibo/app/requests"
  "goweibo/core"
  "goweibo/core/context"
  "goweibo/core/pkg/captcha"
)

type ICaptchaController interface {
  New(*context.AppContext) error
  Verify(*context.AppContext) error
}

type CaptchaController struct {
  URLPrefix string
}

func NewCaptchaController(p string) ICaptchaController {
  return &CaptchaController{URLPrefix: p}
}

func (ca *CaptchaController) New(c *context.AppContext) error {
    return c.AWSuccessJSON(captcha.New(func(id string) string {
      return fmt.Sprintf(
        "%s%s/%s",
        core.GetConfig().String("APP.URL"),
        ca.URLPrefix,
        id,
      )
    }))
}

func (ca *CaptchaController) Verify(c *context.AppContext) error {
  req := new(requests.CaptchaVerify)
  if err := c.AWBindValidatorStruct(req); err != nil {
    return err
  }
  if err := req.Verify(); err != nil {
    return err
  }

  return c.AWSuccessJSON(nil)
}
