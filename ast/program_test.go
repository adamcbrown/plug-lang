package ast_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestProgram(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   ast.Program
	}{
		{
			name:   "basic assignment",
			source: "x = 1\n\ny=2",
			want: ast.Program{
				Assignments: []ast.Assignment{
					{
						Name: ast.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  "x",
								Start: 0,
							},
						},
						Expr: ast.Constant[int]{
							Token: token.Token{
								Type:  token.Integer,
								Text:  "1",
								Start: 4,
							},
							Value: 1,
						},
					},
					{
						Name: ast.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  "y",
								Start: 7,
							},
						},
						Expr: ast.Constant[int]{
							Token: token.Token{
								Type:  token.Integer,
								Text:  "2",
								Start: 9,
							},
							Value: 2,
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer(tc.source))
			got, err := ast.ParseProgram(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
