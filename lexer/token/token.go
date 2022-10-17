package token

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Type int

const (
	Invalid Type = iota
	Identifier
	Integer
	Whitespace
	Character

	// Keywords
	Fn
	EOF
	Unknown
)

var SingleCharacters mapset.Set[rune] = mapset.NewSet(
	'=', '{', '}', '(', ')', '-', '>', ':',
)

var Keywords map[string]Type = map[string]Type{
	"fn": Fn,
}

type Token struct {
	Type  Type
	Text  string
	Start int
}

func (t Token) StartPos() int {
	return t.Start
}

func (t Token) EndPos() int {
	return t.Start + len(t.Text)
}

func (t Token) IsValid() bool {
	return t.Type != Invalid
}

func (t Token) IsRune(r rune) bool {
	return t.Type == Character && len(t.Text) == 1 && t.Text == string(r)
}
