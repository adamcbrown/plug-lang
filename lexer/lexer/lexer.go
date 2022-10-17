package lexer

import (
	"unicode/utf8"

	"github.com/acbrown/plug-lang/lexer/token"
)

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}

type Lexer struct {
	data string
	pos  int
}

func NewLexer(data string) *Lexer {
	return &Lexer{
		data: data,
		pos:  0,
	}
}

func (l *Lexer) current() rune {
	r, _ := utf8.DecodeRuneInString(l.data[l.pos:])
	return r
}

func (l *Lexer) makeToken(typ token.Type, start int) token.Token {
	return token.Token{
		Type:  typ,
		Text:  l.data[start:l.pos],
		Start: start,
	}
}

func (l *Lexer) Lex() token.Token {
	if len(l.data) == l.pos {
		return token.Token{Type: token.EOF, Text: "<EOF>", Start: l.pos}
	}
	start := l.pos
	if token.SingleCharacters.Contains(l.current()) {
		l.pos += 1
		return l.makeToken(token.Character, start)
	}
	if isNumber(l.current()) {
		for l.pos < len(l.data) {
			if !isNumber(l.current()) {
				break
			}
			l.pos += 1
		}
		return l.makeToken(token.Integer, start)
	}
	if isLetter(l.current()) {
		for l.pos < len(l.data) {
			if c := l.current(); !isLetter(c) && !isNumber(c) {
				break
			}
			l.pos += 1
		}
		if kw, ok := token.Keywords[string(l.data[start:l.pos])]; ok {
			return l.makeToken(kw, start)
		}
		return l.makeToken(token.Identifier, start)
	}
	if isWhitespace(l.current()) {
		for l.pos < len(l.data) {
			if !isWhitespace(l.current()) {
				break
			}
			l.pos += 1
		}
		return l.makeToken(token.Whitespace, start)
	}
	l.pos += 1
	return l.makeToken(token.Unknown, start)
}
