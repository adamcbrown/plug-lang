package block

import (
	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/ast/assignment"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Block struct {
	LCurly      token.Token
	Assignments []assignment.Assignment
	RCurly      token.Token
}

var _ ast.Node = Block{}

func (p Block) Start() int {
	return p.LCurly.StartPos()
}

func (p Block) End() int {
	return p.RCurly.EndPos()
}

func Parse(p *parser.Parser) (Block, *ast.ParseErr) {
	lCurly := p.ScanIgnoreWS()
	if lCurly.Type != token.Character || string(lCurly.Text) != "{" {
		return Block{}, &ast.ParseErr{
			Msg: "expected `{` at start of block",
			Tok: lCurly,
		}
	}
	var as []assignment.Assignment
loop:
	for {
		tok := p.ScanIgnoreWS()
		p.Unscan()
		switch tok.Type {
		case token.Identifier:
			a, err := assignment.Parse(p)
			if err != nil {
				return Block{}, err
			}
			as = append(as, a)
		default:
			break loop
		}
	}
	rCurly := p.ScanIgnoreWS()
	if rCurly.Type != token.Character || string(rCurly.Text) != "}" {
		return Block{}, &ast.ParseErr{
			Msg: "expected `}` at start of block",
			Tok: rCurly,
		}
	}
	return Block{
		LCurly:      lCurly,
		Assignments: as,
		RCurly:      rCurly,
	}, nil

}
