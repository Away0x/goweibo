// Package named : 命名路由
package named

import (
	"fmt"
	"gin_weibo/config"
	"gin_weibo/pkg/utils"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	// RouterMap : 存放路由 name 和它的 path map
	RouterMap = make(map[string]string, 0)
	methodMap = make(map[string]string, 0)
)

// Name : 注册路由
// 动态参数目前只支持 :xx 形式
func Name(g gin.IRouter, name string, method string, path string) {
	s := path
	if group, ok := g.(*gin.RouterGroup); ok {
		s = group.BasePath() + path
	}

	if RouterMap[name] != "" {
		panic("该路由已经命名过了: [" + name + "] " + s)
	}
	RouterMap[name] = s
	methodMap[name] = method
}

// G : 根据 name 获取路由 path (完整路径)
// Name(g, "root", "GET", "/")
//     -> G("root") 得到 "/"
// Name(g, "signup.confirm", "POST", "/signup/confirm/:token")
//     -> G("signup.confirm", "token", "abc") 得到 "/signup/confirm/abc"
// Name(g, "users.create", "GET", "/users/create/:id")
//     -> G("users.create", 1) 得到 "/users/create/1"
func G(name string, values ...interface{}) string {
	return config.AppConfig.URL + getRoute(name, values...)
}

// GR : 根据 name 获取路由 path (相对于网站根路径)
func GR(name string, values ...interface{}) string {
	return getRoute(name, values...)
}

// --------- private
func getRoute(name string, values ...interface{}) string {
	path := RouterMap[name]
	valuesArrLen := len(values)

	// 不存在该 name 的路由则 return 一个随机字符串，保证会访问到 404 页面
	if RouterMap[name] == "" {
		return "/" + string(utils.RandomCreateBytes(10))
	}

	// 要么 values length 为 0，要么就为 2
	// values length 为 1
	if valuesArrLen != 0 && valuesArrLen != 1 && valuesArrLen != 2 {
		return path
	}

	// values[0] 为 value，values[0] 的类型为 string 或 int 或 uint
	if valuesArrLen == 1 {
		var val string
		if v, ok := values[0].(uint); ok {
			val = strconv.Itoa(int(v))
		} else if v, ok := values[0].(string); ok {
			val = v
		} else if v, ok := values[0].(int); ok {
			val = strconv.Itoa(v)
		} else {
			return path
		}

		r := strings.NewReplacer(":id", val)
		return r.Replace(path)
	}

	// values[0] 必须为 string，values[1] 必须为 string 或 int(uint)
	if valuesArrLen == 2 {
		var (
			ok  bool
			key string
			val string
		)

		if key, ok = values[0].(string); ok {
			// values[1] 必须为 string 或 int
			if v, ok := values[1].(string); ok {
				val = v
			} else if v, ok := values[1].(int); ok {
				val = strconv.Itoa(v)
			} else if v, ok := values[1].(uint); ok {
				val = strconv.Itoa(int(v))
			} else {
				return path
			}
		}

		r := strings.NewReplacer(":"+key, val)
		return r.Replace(path)
	}

	return path
}

// PrintRoutes 打印 route
func PrintRoutes() {
	for k, v := range RouterMap {
		m := methodMap[k]
		fmt.Fprintf(os.Stderr, "[Route-Name] "+"%-7s %-25s --> %s\n", m, k, v)
	}
}
