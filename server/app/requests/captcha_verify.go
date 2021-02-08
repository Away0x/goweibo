package requests

import (
  "goweibo/core/errno"
  "goweibo/core/pkg/captcha"
  "goweibo/core/pkg/validator"
)

type CaptchaVerify struct {
  Value string `valid:"value"`
  ID    string `valid:"id"`
}

func (c *CaptchaVerify) Options() validator.Options {
  return validator.Options{
    Rules: validator.MapData{
     "value": {"required"},
     "id":    {"required"},
    },
    Messages: validator.MapData{
      "value": {
        "required:验证码不能为空",
      },
      "id": {
        "required:验证码不能为空",
      },
    },
  }
}

func (c *CaptchaVerify) Verify() error {
  if ok := captcha.Verify(c.ID, c.Value); ok {
    return nil
  }
  return errno.ReqErr.WithMessage("验证码错误")
}
