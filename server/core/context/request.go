package context

import (
  "goweibo/core/errno"
  "goweibo/core/pkg/validator"
  "strconv"
)

func (c *AppContext) AWIntParam(key string) (int, error) {
  i, err := strconv.Atoi(c.Param(key))
  if err != nil {
    return 0, err
  }

  return i, nil
}

func (c *AppContext) AWBindValidatorStruct(v validator.Validator) error {
  if err := c.Bind(v); err != nil {
    return err
  }

  errs := validator.ValidateStruct(v)
  if len(errs) > 0 {
    return errno.ReqErr.WitData(errs)
  }

  return nil
}
