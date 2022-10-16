package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
)

type Node interface {
	Start() int
	End() int
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
