package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
)

type Context struct {
	scopes []map[string]Node
}

func (ctx *Context) Resolve(name string) Node {
	for i := range ctx.scopes {
		scope := ctx.scopes[len(ctx.scopes)-1-i]
		if node, ok := scope[name]; ok {
			return node
		}
	}
	return nil
}

func (ctx *Context) PushScope(scope map[string]Node) {
	ctx.scopes = append(ctx.scopes, scope)
}

func (ctx *Context) PopScope() {
	ctx.scopes = ctx.scopes[:len(ctx.scopes)-1]
}

type Node interface {
	Start() int
	End() int
	Enter(ctx *Context)
}

type Expr interface {
	Node
	exprNode()
}

type ExprToken struct{}

func (ExprToken) exprNode() {}

type ParseErr struct {
	Msg string
	Tok token.Token
}
