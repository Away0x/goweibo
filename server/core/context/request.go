package context

import (
  "goweibo/core/errno"
  "goweibo/core/pkg/validator"
  "strconv"
)

func (c *AppContext) AWIntParam(key ...string) (int, error) {
  k := key[0]
  if k == "" {
    k = "id"
  }

  i, err := strconv.Atoi(c.Param(k))
  if err != nil {
    return 0, errno.ReqErr.WithErr(err)
  }

  return i, nil
}

func (c *AppContext) AWIntQuery(key ...string) (int, error) {
  k := key[0]
  if k == "" {
    k = "id"
  }

  i, err := strconv.Atoi(c.QueryParam(k))
  if err != nil {
    return 0, errno.ReqErr.WithErr(err)
  }

  return i, nil
}

func (c *AppContext) AWBindValidatorStruct(v validator.Validator) error {
  if err := c.Bind(v); err != nil {
    return err
  }

  if errs, ok := validator.ValidateStruct(v); !ok {
    return errno.ReqErr.WitData(errs)
  }

  return nil
}
