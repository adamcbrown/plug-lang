package ast_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestAssignment(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   ast.Assignment
	}{
		{
			name:   "basic assignment",
			source: "x = 10",
			want: ast.Assignment{
				Name: ast.Name{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "x",
						Start: 0,
					},
				},
				Expr: &ast.Constant[int]{
					Token: token.Token{
						Type:  token.Integer,
						Text:  "10",
						Start: 4,
					},
					Value: 10,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer(tc.source))
			got, err := ast.ParseAssignment(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
