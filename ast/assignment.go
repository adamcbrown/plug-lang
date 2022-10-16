package ast

import (
	"fmt"

	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Assignment struct {
	Name Name
	Expr Expr
}

var _ Node = Assignment{}

func (a Assignment) Start() int {
	return a.Name.Start()
}

func (a Assignment) End() int {
	return a.Expr.End()
}

func ParseAssignment(p *parser.Parser) (Assignment, *ParseErr) {
	tok := p.Scan()
	if tok.Type != token.Identifier {
		return Assignment{}, &ParseErr{
			Msg: fmt.Sprintf("expected `Identifier` at start of assignment, found %v", tok),
			Tok: tok,
		}
	}
	id := Name{Token: tok}

	if tok := p.ScanIgnoreWS(); tok.Type != token.Character && string(tok.Text) != "=" {
		return Assignment{}, &ParseErr{
			Msg: fmt.Sprintf("expected `=` assignment operator after ID, found %v", tok),
			Tok: tok,
		}
	}

	expr, err := ParseExpr(p)
	if err != nil {
		return Assignment{}, err
	}

	return Assignment{
		Name: id,
		Expr: expr,
	}, nil
}
