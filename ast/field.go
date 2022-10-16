package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Field struct {
	Name Name
	Type Expr
}

var _ Node = Field{}

func (f Field) Start() int {
	if f.Name.Token.IsValid() {
		return f.Name.Start()
	}
	return f.Type.Start()
}

func (f Field) End() int {
	return f.Type.End()
}

func ParseField(p *parser.Parser) (Field, *ParseErr) {
	nameTok := p.ScanIgnoreWS()
	if nameTok.Type != token.Identifier {
		p.Unscan()
		typ, err := ParseExpr(p)
		if err != nil {
			return Field{}, &ParseErr{
				Msg: "expected Identifier or Type at start of field",
				Tok: nameTok,
			}
		}

		return Field{
			Type: typ,
		}, nil
	}

	if tok := p.Scan(); !tok.IsRune(':') {
		p.Unscan()
		p.Unscan()
		typ, err := ParseExpr(p)
		if err != nil {
			return Field{}, &ParseErr{
				Msg: "expected `:` after field identifier",
				Tok: tok,
			}
		}

		return Field{
			Type: typ,
		}, nil
	}

	typ, err := ParseExpr(p)
	if err != nil {
		return Field{}, err
	}

	return Field{
		Name: Name{
			Token: nameTok,
		},
		Type: typ,
	}, nil
}
