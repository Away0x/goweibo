package tpl

import "github.com/flosch/pongo2"

type tagStaticTag struct {
  path string
}

func (node *tagStaticTag) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
  staticPath := node.path

  if string(staticPath[0]) == "/" {
    writer.WriteString(config.GeneratePublicPath(staticPath))
  } else {
    writer.WriteString(config.GeneratePublicPath("/" + staticPath))
  }
  return nil
}

func StaticTag(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
  pathToken := arguments.MatchType(pongo2.TokenString)
  if pathToken == nil {
    return nil, arguments.Error("static tag error: path 必须为 string.", nil)
  }

  node := &tagStaticTag{path: pathToken.Val}

  if arguments.Remaining() > 0 {
    return nil, arguments.Error("Malformed static-tag arguments.", nil)
  }

  return node, nil
}
