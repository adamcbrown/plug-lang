package program

import (
	"testing"

	"github.com/acbrown/plug-lang/ast/assignment"
	"github.com/acbrown/plug-lang/ast/expr"
	"github.com/acbrown/plug-lang/ast/name"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestProgram(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   Program
	}{
		{
			name:   "basic assignment",
			source: "x = 1\n\ny=2",
			want: Program{
				Assignments: []assignment.Assignment{
					{
						Name: name.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("x"),
								Start: 0,
							},
						},
						Expr: expr.Constant[int]{
							Token: token.Token{
								Type:  token.Integer,
								Text:  []rune("1"),
								Start: 4,
							},
							Value: 1,
						},
					},
					{
						Name: name.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("y"),
								Start: 7,
							},
						},
						Expr: expr.Constant[int]{
							Token: token.Token{
								Type:  token.Integer,
								Text:  []rune("2"),
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
			p := parser.NewParser(lexer.NewLexer([]rune(tc.source)))
			got, err := Parse(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}