package ast

type Modification struct {
	ExprToken
	Base  Expr
	Block Block
}

var _ Expr = Modification{}

func (m Modification) Start() int {
	return m.Base.Start()
}

func (m Modification) End() int {
	return m.Block.End()
}
