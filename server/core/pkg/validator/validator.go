package validator

import (
  "github.com/thedevsaddam/govalidator"
)

type Validator interface {
  Options() Options
}

type MapData = govalidator.MapData

type Options struct {
  TagIdentifier string
  Rules MapData
  Messages MapData
}

type BaseValidator struct {}

func (v *BaseValidator) Options() Options {
  return Options{}
}

func ValidateStruct(v Validator) (map[string][]string, bool) {
  o := v.Options()
  if o.TagIdentifier == "" {
    o.TagIdentifier = "valid"
  }

  errs := govalidator.New(govalidator.Options{
    TagIdentifier: o.TagIdentifier,
    Data: v,
    Rules: o.Rules,
    Messages: o.Messages,
  }).ValidateStruct()

  if len(errs) > 0 {
    return errs, false
  }
  return nil, true
}
