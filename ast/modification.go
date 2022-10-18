package ast

import "github.com/acbrown/plug-lang/types"

type Modification struct {
	ExprToken
	Base  Expr
	Block Block
}

var _ Expr = &Modification{}

func (m *Modification) Start() int {
	return m.Base.Start()
}

func (m *Modification) End() int {
	return m.Block.End()
}

func (m *Modification) AddReferences(ctx *Context) {
	m.Base.AddReferences(ctx)
	m.Block.AddReferences(ctx)
}

func (m *Modification) Type(ctx *Context) types.Type {
	return m.Base.AsType(ctx)
}

func (m *Modification) AsType(ctx *Context) types.Type {
	ctx.AddError(NodeErr{
		Msg: "modification is not a type",
		N:   m,
	})
	return types.ErrorType
}
