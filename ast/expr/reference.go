package expr

import (
	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/token"
)

type Reference struct {
	ExprToken
	Token token.Token
}

var _ ast.Node = Reference{}
var _ Expr = Reference{}

func (r Reference) Start() int {
	return r.Token.StartPos()
}

func (r Reference) End() int {
	return r.Token.EndPos()
}
