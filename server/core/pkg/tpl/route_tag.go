package tpl

import "github.com/flosch/pongo2"

type tagRouteTag struct {
	routerName string
	args       []pongo2.IEvaluator
}

func (node *tagRouteTag) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	args := make([]interface{}, 0)
	routerName := node.routerName

	for _, arg := range node.args {
		val, err := arg.Evaluate(ctx)
		if err != nil {
			return err
		}

		args = append(args, val.Interface())
	}

	routePath := config.GetRoutePath(routerName, args...)
	writer.WriteString(routePath)
	return nil
}

// RouteTag 获取到路由
func RouteTag(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	nameToken := arguments.MatchType(pongo2.TokenString)
	if nameToken == nil {
		return nil, arguments.Error("route tag error: name 必须为 string.", nil)
	}

	routeTag := &tagRouteTag{
		routerName: nameToken.Val,
	}

	for arguments.Remaining() > 0 {
		node, err := arguments.ParseExpression()
		if err != nil {
			return nil, err
		}
		routeTag.args = append(routeTag.args, node)
	}

	return routeTag, nil
}
