package name

import (
	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/token"
)

type Name struct {
	Token token.Token
}

var _ ast.Node = Name{}

func (i Name) Start() int {
	return i.Token.StartPos()
}

func (i Name) End() int {
	return i.Token.EndPos()
}
