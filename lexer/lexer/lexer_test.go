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
				{Type: token.Integer, Text: "1", Start: 0},
				{Type: token.Whitespace, Text: " ", Start: 1},
				{Type: token.Integer, Text: "02", Start: 2},
				{Type: token.Whitespace, Text: "\t", Start: 4},
				{Type: token.Integer, Text: "12", Start: 5},
				{Type: token.Whitespace, Text: "\n", Start: 7},
				{Type: token.EOF, Text: "<EOF>", Start: 8},
			},
		},
		{
			name:  "Identifier",
			input: "Hello",
			want: []token.Token{
				{Type: token.Identifier, Text: "Hello", Start: 0},
				{Type: token.EOF, Text: "<EOF>", Start: 5},
			},
		},
		{
			name:  "Identifier with letter",
			input: "Hello123 x",
			want: []token.Token{
				{Type: token.Identifier, Text: "Hello123", Start: 0},
				{Type: token.Whitespace, Text: " ", Start: 8},
				{Type: token.Identifier, Text: "x", Start: 9},
				{Type: token.EOF, Text: "<EOF>", Start: 10},
			},
		},
		{
			name:  "Assign no space",
			input: "x=1",
			want: []token.Token{
				{Type: token.Identifier, Text: "x", Start: 0},
				{Type: token.Character, Text: "=", Start: 1},
				{Type: token.Integer, Text: "1", Start: 2},
				{Type: token.EOF, Text: "<EOF>", Start: 3},
			},
		},
		{
			name:  "Assign with space",
			input: "x = 1",
			want: []token.Token{
				{Type: token.Identifier, Text: "x", Start: 0},
				{Type: token.Whitespace, Text: " ", Start: 1},
				{Type: token.Character, Text: "=", Start: 2},
				{Type: token.Whitespace, Text: " ", Start: 3},
				{Type: token.Integer, Text: "1", Start: 4},
				{Type: token.EOF, Text: "<EOF>", Start: 5},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			lexer := lexer.NewLexer(tc.input)

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
