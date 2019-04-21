package requests

import (
	"gin_weibo/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

type (
	// 验证器函数
	ValidatorFunc = func() (msg string)
	// 验证器数组 map
	ValidatorMap = map[string][]ValidatorFunc
	// 错误信息数组
	ValidatorMsgArr = map[string][]string
)

/*
验证器数组按顺序验证，一旦验证没通过，即结束该字段的验证

RunValidators(
  ValidatorMap{
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
  ValidatorMsgArr{
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
func RunValidators(m ValidatorMap, msgMap ValidatorMsgArr) (errors []string) {
	// start := time.Now()
	// defer func() {
	// 	cost := time.Since(start)
	// 	fmt.Println("cost=========", cost)
	// }()

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

// func RunValidators(m ValidatorMap, msgMap ValidatorMsgArr) (errors []string) {
// 	start := time.Now()
// 	defer func() {
// 		cost := time.Since(start)
// 		fmt.Println("cost=========", cost)
// 	}()

// 	wg := sync.WaitGroup{}
// 	finished := make(chan bool, 1)

// 	for key, validators := range m {
// 		wg.Add(1)
// 		customMsgArr := msgMap[key] // 自定义错误信息数组
// 		customMsgArrLen := len(customMsgArr)

// 		go func(k string, v []ValidatorFunc) {
// 			defer wg.Done()

// 			for i, fn := range v {
// 				msg := fn()
// 				if msg != "" {
// 					if i < customMsgArrLen && customMsgArr[i] != "" {
// 						// 采用自定义的错误信息输出
// 						msg = customMsgArr[i]
// 					} else {
// 						// 采用默认的错误信息输出
// 						names := strings.Split(k, "|")
// 						data := make(map[string]string)

// 						for ti, tv := range names {
// 							data["$key"+strconv.Itoa(ti+1)+"$"] = tv
// 						}

// 						msg = utils.ParseEasyTemplate(msg, data)
// 					}

// 					errors = append(errors, msg)
// 					break // 进行下一个字段的验证
// 				}
// 			}
// 		}(key, validators)
// 	}

// 	go func() {
// 		wg.Wait() // 上面多个 goroutine 的并行处理完会发送消息给 finished
// 		close(finished)
// 	}()

// 	// 等待消息 (无可用 case 也无 default 会堵塞)
// 	select {
// 	case <-finished:
// 	}

// 	return errors
// }

// RequiredValidator : value 必须存在
func RequiredValidator(value string) ValidatorFunc {
	return func() (msg string) {
		if value == "" {
			return "$key1$ 必须存在"
		}

		return ""
	}
}

// MixLengthValidator -
func MixLengthValidator(value string, minStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)

		if l < minStrLen {
			return "$key1$ 必须大于 " + strconv.Itoa(minStrLen)
		}

		return ""
	}
}

// MaxLengthValidator -
func MaxLengthValidator(value string, maxStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)

		if l > maxStrLen {
			return "$key1$ 必须小于 " + strconv.Itoa(maxStrLen)
		}

		return ""
	}
}

// EqualValidator -
func EqualValidator(v1 string, v2 string) ValidatorFunc {
	return func() (msg string) {
		if v1 != v2 {
			return "$key1$ 必须等于 $key2$"
		}

		return ""
	}
}

// EmailValidator 验证邮箱格式
func EmailValidator(value string) ValidatorFunc {
	return func() (msg string) {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		status := reg.MatchString(value)

		if !status {
			return "$key1$ 邮箱格式错误"
		}

		return ""
	}
}
