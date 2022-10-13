package lexer_test

import (
	"testing"

	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/google/go-cmp/cmp"
)

func TestLexer(t *testing.T) {
	tcs := []struct {
		name  string
		input string
		want  []token.Token
	}{
		{
			name:  "Integers",
			input: "1 02\t12\n",
			want: []token.Token{
				{Type: token.Integer, Text: []rune("1"), Start: 0},
				{Type: token.Whitespace, Text: []rune(" "), Start: 1},
				{Type: token.Integer, Text: []rune("02"), Start: 2},
				{Type: token.Whitespace, Text: []rune("\t"), Start: 4},
				{Type: token.Integer, Text: []rune("12"), Start: 5},
				{Type: token.Whitespace, Text: []rune("\n"), Start: 7},
				{Type: token.EOF, Text: []rune("<EOF>"), Start: 8},
			},
		},
		{
			name:  "Identifier",
			input: "Hello",
			want: []token.Token{
				{Type: token.Identifier, Text: []rune("Hello"), Start: 0},
				{Type: token.EOF, Text: []rune("<EOF>"), Start: 5},
			},
		},
		{
			name:  "Identifier with letter",
			input: "Hello123 x",
			want: []token.Token{
				{Type: token.Identifier, Text: []rune("Hello123"), Start: 0},
				{Type: token.Whitespace, Text: []rune(" "), Start: 8},
				{Type: token.Identifier, Text: []rune("x"), Start: 9},
				{Type: token.EOF, Text: []rune("<EOF>"), Start: 10},
			},
		},
		{
			name:  "Assign no space",
			input: "x=1",
			want: []token.Token{
				{Type: token.Identifier, Text: []rune("x"), Start: 0},
				{Type: token.Character, Text: []rune("="), Start: 1},
				{Type: token.Integer, Text: []rune("1"), Start: 2},
				{Type: token.EOF, Text: []rune("<EOF>"), Start: 3},
			},
		},
		{
			name:  "Assign with space",
			input: "x = 1",
			want: []token.Token{
				{Type: token.Identifier, Text: []rune("x"), Start: 0},
				{Type: token.Whitespace, Text: []rune(" "), Start: 1},
				{Type: token.Character, Text: []rune("="), Start: 2},
				{Type: token.Whitespace, Text: []rune(" "), Start: 3},
				{Type: token.Integer, Text: []rune("1"), Start: 4},
				{Type: token.EOF, Text: []rune("<EOF>"), Start: 5},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			lexer := lexer.NewLexer([]rune(tc.input))

			var tokens []token.Token
			for {
				tok := lexer.Lex()
				tokens = append(tokens, tok)
				if tok.Type == token.EOF {
					break
				}
			}

			if diff := cmp.Diff(tc.want, tokens); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
