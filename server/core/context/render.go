package context

import "net/http"

// TplData 模版数据类型
type TplData map[string]interface{}

// RenderHTML 渲染 html
func (c *AppContext) RenderHTML(tpl string, data interface{}) error {
	tplname := tpl + ".tpl"

	if typed, ok := data.(TplData); ok {
		return c.Render(http.StatusOK, tplname, map[string]interface{}(typed))
	}

	return c.Render(http.StatusOK, tplname, data)
}

// RenderHTMLNoData 渲染无数据的 html
func (c *AppContext) RenderHTMLNoData(tpl string) error {
	return c.RenderHTML(tpl, map[string]interface{}{})
}
