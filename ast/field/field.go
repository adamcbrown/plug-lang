package field

import (
	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/ast/expr"
	"github.com/acbrown/plug-lang/ast/name"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Field struct {
	Name name.Name
	Type expr.Expr
}

var _ ast.Node = Field{}

func (f Field) Start() int {
	if f.Name.Token.IsValid() {
		return f.Name.Start()
	}
	return f.Type.Start()
}

func (f Field) End() int {
	return f.Type.End()
}

func Parse(p *parser.Parser) (Field, *ast.ParseErr) {
	nameTok := p.ScanIgnoreWS()
	if nameTok.Type != token.Identifier {
		p.Unscan()
		typ, err := expr.Parse(p)
		if err != nil {
			return Field{}, &ast.ParseErr{
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
		typ, err := expr.Parse(p)
		if err != nil {
			return Field{}, &ast.ParseErr{
				Msg: "expected `:` after field identifier",
				Tok: tok,
			}
		}

		return Field{
			Type: typ,
		}, nil
	}

	typ, err := expr.Parse(p)
	if err != nil {
		return Field{}, err
	}

	return Field{
		Name: name.Name{
			Token: nameTok,
		},
		Type: typ,
	}, nil
}
