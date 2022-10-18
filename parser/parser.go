package parser

import (
	"log"

	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
)

const BUFFER_MAX = 2

type Parser struct {
	l *lexer.Lexer

	typeParseCounter int

	buffer    []token.Token
	bufferIdx int
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		l:      l,
		buffer: make([]token.Token, 0, BUFFER_MAX),
	}
}

func (p *Parser) EnterType() {
	p.typeParseCounter += 1
}

func (p *Parser) ExitType() {
	p.typeParseCounter -= 1
}

func (p *Parser) InFunction() bool {
	return p.typeParseCounter > 0
}

func (p *Parser) Scan() token.Token {
	if p.bufferIdx < len(p.buffer) {
		tok := p.buffer[p.bufferIdx]
		p.bufferIdx += 1
		return tok
	}

	tok := p.l.Lex()
	if len(p.buffer) < BUFFER_MAX {
		p.buffer = append(p.buffer, tok)
		p.bufferIdx += 1
		return tok
	}

	// Shift elements in buffer down 1 element
	copy(p.buffer, p.buffer[1:])
	p.buffer[BUFFER_MAX-1] = tok
	return tok
}

func (p *Parser) ScanIgnoreWS() token.Token {
	if tok := p.Scan(); tok.Type != token.Whitespace {
		return tok
	}
	return p.Scan()
}

func (p *Parser) Unscan() {
	if p.bufferIdx == 0 {
		log.Fatal("Went below buffer. Need to increase buffer limit")
	}
	p.bufferIdx -= 1
}
