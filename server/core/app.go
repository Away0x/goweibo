package core

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/labstack/echo/v4"
)

// Application global application
type Application struct {
	Engine *echo.Echo
}

// NewApplication setup application
func NewApplication(e *echo.Echo) {
	application = &Application{
		Engine: e,
	}
}

// RoutePath 根据 route name 获取 route path
func (a *Application) RoutePath(name string, params ...interface{}) string {
	return a.Engine.Reverse(name, params...)
}

// PrintRoutes 输出路由配置
func (a *Application) PrintRoutes(filename string) {
	routes := make([]*echo.Route, 0)
	for _, item := range a.Engine.Routes() {
		if strings.HasPrefix(item.Name, "github.com") {
			continue
		}

		routes = append(routes, item)
	}

	routesStr, _ := json.MarshalIndent(struct {
		Count  int           `json:"count"`
		Routes []*echo.Route `json:"routes"`
	}{
		Count:  len(routes),
		Routes: routes,
	}, "", "  ")

	ioutil.WriteFile(filename, routesStr, 0644)
}
