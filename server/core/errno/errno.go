package errno

import (
  "fmt"
  "goweibo/core/constants"
)

// Errno 项目统一错误类型
type Errno struct {
  HTTPCode int                 // http 状态码
  Message  string              // 错误信息
  Code     constants.LogicCode // 自定义错误码
  Err      error               // 具体错误
  Data     interface{}         // 承载一些具体的错误内容
}

func (e Errno) Error() string {
  return e.Message
}

// WithErr 复写 err
func (e Errno) WithErr(err error) error {
  if err == nil {
    return nil
  }

  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  err.Error(),
    Err:      err,
  }
}

// WitData 复写 err
func (e Errno) WitData(d interface{}) error {
  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Data:     d,
  }
}

// WithMessage 复写 message
func (e Errno) WithMessage(msg string) error {
  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  msg,
    Err:      e.Err,
  }
}

// WithMessagef 复写 message
func (e Errno) WithMessagef(format string, args ...interface{}) error {
  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  fmt.Sprintf(format, args...),
    Err:      e.Err,
  }
}

// WithErrMessage 复写 err 和 message
func (e Errno) WithErrMessage(err error, msg string) error {
  if err == nil {
    return nil
  }

  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  msg,
    Err:      err,
  }
}

// WithErrMessagef 复写 err 和 message
func (e Errno) WithErrMessagef(err error, format string, args ...interface{}) error {
  if err == nil {
    return nil
  }

  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  fmt.Sprintf(format, args...),
    Err:      err,
  }
}
