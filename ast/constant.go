package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
)

type Constant[T any] struct {
	ExprToken
	Token token.Token
	Value T
}

var _ Node = Constant[any]{}
var _ Expr = Constant[any]{}

func (c Constant[T]) Start() int {
	return c.Token.StartPos()
}

func (c Constant[T]) End() int {
	return c.Token.EndPos()
}

func (c Constant[T]) Enter(ctx *Context)
