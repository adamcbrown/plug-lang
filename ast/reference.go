package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
)

type Reference struct {
	ExprToken
	Token token.Token
}

var _ Node = Reference{}
var _ Expr = Reference{}

func (r Reference) Start() int {
	return r.Token.StartPos()
}

func (r Reference) End() int {
	return r.Token.EndPos()
}

func (r Reference) Enter(ctx *Context) {}
