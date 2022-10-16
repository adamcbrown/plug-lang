package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Block struct {
	LCurly      token.Token
	Assignments []Assignment
	RCurly      token.Token
}

var _ Node = Block{}

func (p Block) Start() int {
	return p.LCurly.StartPos()
}

func (p Block) End() int {
	return p.RCurly.EndPos()
}

func ParseBlock(p *parser.Parser) (Block, *ParseErr) {
	lCurly := p.ScanIgnoreWS()
	if lCurly.Type != token.Character || string(lCurly.Text) != "{" {
		return Block{}, &ParseErr{
			Msg: "expected `{` at start of block",
			Tok: lCurly,
		}
	}
	var as []Assignment
loop:
	for {
		tok := p.ScanIgnoreWS()
		p.Unscan()
		switch tok.Type {
		case token.Identifier:
			a, err := ParseAssignment(p)
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
		return Block{}, &ParseErr{
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
