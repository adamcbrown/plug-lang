package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Program struct {
	Assignments []Assignment
}

var _ Node = Program{}

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
		default:
			return Program{as}, nil
		}
	}

}
