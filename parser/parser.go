package parser

import (
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
)

type Parser struct {
	l *lexer.Lexer

	tok      token.Token
	buffered bool
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		l: l,
	}
}

func (p *Parser) Scan() token.Token {
	if p.buffered {
		p.buffered = false
		return p.tok
	}

	p.tok = p.l.Lex()
	return p.tok
}

func (p *Parser) ScanIgnoreWS() token.Token {
	if tok := p.Scan(); tok.Type != token.Whitespace {
		return tok
	}
	return p.Scan()
}

func (p *Parser) Unscan() {
	p.buffered = true
}
