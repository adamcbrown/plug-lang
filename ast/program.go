package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Program struct {
	Assignments []Assignment
}

var _ Node = &Program{}

func (p *Program) Start() int {
	if len(p.Assignments) == 0 {
		return 0
	}
	return p.Assignments[0].Start()
}

func (p *Program) End() int {
	if len(p.Assignments) == 0 {
		return 0
	}
	return p.Assignments[len(p.Assignments)-1].End()
}

func (p *Program) AddReferences(ctx *Context) {
	scope := make(map[string]Expr, len(p.Assignments))
	for i := range p.Assignments {
		a := &p.Assignments[i]
		scope[a.Name.Token.Text] = a.Expr
	}

	ctx.PushScope(scope)
	defer ctx.PopScope()

	for i := range p.Assignments {
		(&p.Assignments[i]).AddReferences(ctx)
	}
}

func ParseProgram(p *parser.Parser) (Program, *ParseErr) {
	var as []Assignment
	for {
		tok := p.ScanIgnoreWS()
		p.Unscan()
		switch tok.Type {
		case token.Identifier:
			a, err := ParseAssignment(p)
			if err != nil {
				return Program{}, err
			}
			as = append(as, a)
		case token.EOF:
			return Program{as}, nil
		default:
			return Program{}, &ParseErr{
				Msg: "unknown token",
				Tok: tok,
			}
		}
	}

}
