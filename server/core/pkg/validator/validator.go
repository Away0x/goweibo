package validator

import (
  "github.com/thedevsaddam/govalidator"
)

type Validator interface {
  Options() govalidator.Options
}

type BaseValidator struct {}

func (v *BaseValidator) DefaultTagIdentifier() string {
  return "json"
}

func (v *BaseValidator) Options() govalidator.Options {
  return govalidator.Options{}
}

func ValidateStruct(v Validator) map[string][]string {
  errs := govalidator.New(v.Options()).ValidateStruct()
  return errs
}
