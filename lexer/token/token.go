package token

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Type int

const (
	Identifier Type = iota
	Integer
	Whitespace
	Character
	EOF
	Unknown
)

var SingleCharacters mapset.Set[rune] = mapset.NewSet(
	'=',
)

type Token struct {
	Type  Type
	Text  []rune
	Start int
}

func (t Token) StartPos() int {
	return t.Start
}

func (t Token) EndPos() int {
	return t.Start + len(t.Text)
}
