package context

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// AppContext 项目的 context
type AppContext struct {
	echo.Context
}

type (
	// AppHandlerFunc 项目 handler 定义
	AppHandlerFunc = func(c *AppContext) error
	// EchoRegisterFunc echo handler 类型
	EchoRegisterFunc = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
)

// NewAppContext new AppContext
func NewAppContext(c echo.Context) *AppContext {
	return &AppContext{Context: c}
}

// RegisterHandler 注册路由
func RegisterHandler(fn EchoRegisterFunc, path string, h AppHandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return fn(path, func(c echo.Context) error {
		cc, ok := c.(*AppContext)
		if !ok {
			cc = NewAppContext(c)
			return h(cc)
		}
		return h(cc)
	}, m...)
}
