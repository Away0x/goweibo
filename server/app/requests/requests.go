package requests

import (
  "errors"
  "fmt"
  "github.com/thedevsaddam/govalidator"
  "goweibo/app/models"
  "strconv"
  "strings"
  "unicode/utf8"
)

func init() {
  // not_exists:users
  govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
    rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

    tableName := rng[0]
    dbFiled := field
    if len(rng) == 2 && rng[1] != "" {
      dbFiled = rng[1]
    }
    val := value.(string)

    var count int64
    models.DB().Table(tableName).Where(dbFiled+" = ?", val).Count(&count)

    if count != 0 {
      if message != "" {
        return errors.New(message)
      }

      return fmt.Errorf("%v 已被占用", val)
    }
    return nil
  })

  // max_cn:8
  govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
    valLength := utf8.RuneCountInString(value.(string))
    l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:")) // handle other error
    if valLength > l {
      if message != "" {
        return errors.New(message)
      }
      return fmt.Errorf("长度不能超过 %d 个字", l)
    }
    return nil
  })

  // min_cn:2
  govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
    valLength := utf8.RuneCountInString(value.(string))
    l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:")) // handle other error
    if valLength < l {
      if message != "" {
        return errors.New(message)
      }
      return fmt.Errorf("长度需大于 %d 个字", l)
    }
    return nil
  })
}
