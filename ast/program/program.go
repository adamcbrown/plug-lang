package program

import (
	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/ast/assignment"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Program struct {
	Assignments []assignment.Assignment
}

var _ ast.Node = Program{}

func (p Program) Start() int {
	if len(p.Assignments) == 0 {
		return 0
	}
	return p.Assignments[0].Start()
}

func (p Program) End() int {
	if len(p.Assignments) == 0 {
		return 0
	}
	return p.Assignments[len(p.Assignments)-1].End()
}

func Parse(p *parser.Parser) (Program, *ast.ParseErr) {
	var as []assignment.Assignment
	for {
		tok := p.ScanIgnoreWS()
		p.Unscan()
		switch tok.Type {
		case token.Identifier:
			a, err := assignment.Parse(p)
			if err != nil {
				return Program{}, err
			}
			as = append(as, a)
		default:
			return Program{as}, nil
		}
	}

}
