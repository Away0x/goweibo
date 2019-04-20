package requests

import (
	"gin_weibo/pkg/utils"
	"strconv"
	"strings"
)

type (
	// 验证器函数
	validatorFunc = func() (msg string)
	// 验证器数组 map
	validatorMap = map[string][]validatorFunc
	// 错误信息数组
	validatorMsgArr = map[string][]string
)

/*
验证器数组按顺序验证，一旦验证没通过，即结束该字段的验证

RunValidators(
  validatorMap{
    "name": {
      RequiredValidator(u.Name),
      MaxLengthValidator(u.Name, 50),
    },
    "email": {
      RequiredValidator(u.Email),
      MaxLengthValidator(u.Email, 255),
      u.uniqueValidator(),
    },
    "password|password_confirmation": { // "key1|key2" 语法会注入错误信息中，替换 $key1$ 和 $key2$ (用了这种用法，就不能使用自定义错误信息了)
      RequiredValidator(u.Password),
      MixLengthValidator(u.Password, 6),
      EqualValidator(u.Password, u.PasswordConfirmation),
    },
  },
  validatorMsgArr{
    "name": {
      "名称不能为空", // 自定义错误信息 (需按验证器注册顺序摆放, "" 表示使用默认错误信息)
    },
  },
)

以上调用，如果验证失败会产生如下输出:
[
  名称不能为空
  email 必须存在
  password 必须大于 6
  password 必须等于 password_confirmation
]
*/
func RunValidators(m validatorMap, msgMap validatorMsgArr) (errors []string) {
	for k, validators := range m {
		customMsgArr := msgMap[k] // 自定义错误信息数组
		customMsgArrLen := len(customMsgArr)

		for i, fn := range validators {
			msg := fn()
			if msg != "" {
				if i < customMsgArrLen && customMsgArr[i] != "" {
					// 采用自定义的错误信息输出
					msg = customMsgArr[i]
				} else {
					// 采用默认的错误信息输出
					names := strings.Split(k, "|")
					data := make(map[string]string)

					for ti, tv := range names {
						data["$key"+strconv.Itoa(ti+1)+"$"] = tv
					}

					msg = utils.ParseEasyTemplate(msg, data)
				}

				errors = append(errors, msg)
				break // 进行下一个字段的验证
			}
		}
	}

	return errors
}

// RequiredValidator : value 必须存在
func RequiredValidator(value string) validatorFunc {
	return func() (msg string) {
		if value == "" {
			return "$key1$ 必须存在"
		}

		return ""
	}
}

// MixLengthValidator -
func MixLengthValidator(value string, minStrLen int) validatorFunc {
	return func() (msg string) {
		l := len(value)

		if l < minStrLen {
			return "$key1$ 必须大于 " + strconv.Itoa(minStrLen)
		}

		return ""
	}
}

// MaxLengthValidator -
func MaxLengthValidator(value string, maxStrLen int) validatorFunc {
	return func() (msg string) {
		l := len(value)

		if l > maxStrLen {
			return "$key1$ 必须小于 " + strconv.Itoa(maxStrLen)
		}

		return ""
	}
}

// EqualValidator -
func EqualValidator(v1 string, v2 string) validatorFunc {
	return func() (msg string) {
		if v1 != v2 {
			return "$key1$ 必须等于 $key2$"
		}

		return ""
	}
}
