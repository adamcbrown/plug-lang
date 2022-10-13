package assignment

import (
	"fmt"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/ast/expr"
	"github.com/acbrown/plug-lang/ast/name"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Assignment struct {
	Name name.Name
	Expr expr.Expr
}

var _ ast.Node = Assignment{}

func (a Assignment) Start() int {
	return a.Name.Start()
}

func (a Assignment) End() int {
	return a.Expr.End()
}

func Parse(p *parser.Parser) (Assignment, *ast.ParseErr) {
	tok := p.Scan()
	if tok.Type != token.Identifier {
		return Assignment{}, &ast.ParseErr{
			Msg: fmt.Sprintf("expected `Identifier` at start of assignment, found %v", tok),
			Tok: tok,
		}
	}
	id := name.Name{Token: tok}

	if tok := p.ScanIgnoreWS(); tok.Type != token.Character && string(tok.Text) != "=" {
		return Assignment{}, &ast.ParseErr{
			Msg: fmt.Sprintf("expected `=` assignment operator after ID, found %v", tok),
			Tok: tok,
		}
	}

	expr, err := expr.Parse(p)
	if err != nil {
		return Assignment{}, err
	}

	return Assignment{
		Name: id,
		Expr: expr,
	}, nil
}
