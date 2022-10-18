package ast

import (
	"log"

	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/types"
)

type Constant[T any] struct {
	ExprToken
	Token token.Token
	Value T
}

var _ Expr = &Constant[any]{}

func (c *Constant[T]) Start() int {
	return c.Token.StartPos()
}

func (c *Constant[T]) End() int {
	return c.Token.EndPos()
}

func (c *Constant[T]) AddReferences(ctx *Context) {}

func (c *Constant[T]) Type(*Context) types.Type {
	switch c.Token.Type {
	case token.Integer:
		return types.IntType
	}
	log.Fatalf("Unsupported Constant Type %T", c.Value)
	return nil
}

func (c *Constant[T]) AsType(ctx *Context) types.Type {
	ctx.AddError(NodeErr{
		Msg: "modification is not a type",
		N:   c,
	})
	return types.ErrorType
}
