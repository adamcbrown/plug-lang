package ast

import "github.com/acbrown/plug-lang/lexer/token"

type Node interface {
	Start() int
	End() int
}

type ParseErr struct {
	Msg string
	Tok token.Token
}
