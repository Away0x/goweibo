package captcha

import (
  "github.com/dchest/captcha"
  "github.com/labstack/echo/v4"
)

type Captcha struct {
  ID       string
  ImageURL string
  AudioURL string
}

// New 获取验证码
func New(getURL func(id string) string) Captcha {
  id := captcha.New()
  url := getURL(id)

  return Captcha{
    ID:       id,
    ImageURL: url + ".png",
    AudioURL: url + ".wav",
  }
}

// Verify 验证
func Verify(captchaID, captchaVal string) bool {
  return captcha.VerifyString(captchaID, captchaVal)
}

// Handler gin handler
func Handler(c echo.Context) error {
  ServeHTTP(c.Response().Writer, c.Request())
  return nil
}
