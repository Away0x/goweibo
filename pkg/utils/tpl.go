package utils

import "strings"

// ParseEasyTemplate : 简单易用的解析模板方法
func ParseEasyTemplate(tplString string, data map[string]string) string {
	replaceArr := []string{}
	for k, v := range data {
		replaceArr = append(replaceArr, k, v)
	}

	r := strings.NewReplacer(replaceArr...)

	return r.Replace(tplString)
}
