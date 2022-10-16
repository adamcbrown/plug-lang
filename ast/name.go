package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
)

type Name struct {
	Token token.Token
}

var _ Node = Name{}

func (i Name) Start() int {
	return i.Token.StartPos()
}

func (i Name) End() int {
	return i.Token.EndPos()
}
