package context

import (
  "net/http"

	"github.com/spf13/viper"
)

// TplData 模版数据类型
type TplData map[string]interface{}

// AWHtml 渲染 html
func (c *AppContext) AWHtml(tpl string, data interface{}) error {
	tplname := tpl + "." + viper.GetString("APP.TEMPLATE_EXT")

	if typed, ok := data.(TplData); ok {
		return c.Render(http.StatusOK, tplname, map[string]interface{}(typed))
	}

	return c.Render(http.StatusOK, tplname, data)
}

// AWHtmlNoData 渲染无数据的 html
func (c *AppContext) AWHtmlNoData(tpl string) error {
	return c.AWHtml(tpl, map[string]interface{}{})
}
