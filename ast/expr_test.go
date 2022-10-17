package ast_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestAssignmentExpr(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   ast.Expr
	}{
		{
			name:   "integer",
			source: "10",
			want: ast.Constant[int]{
				Token: token.Token{
					Type:  token.Integer,
					Text:  "10",
					Start: 0,
				},
				Value: 10,
			},
		},
		{
			name:   "reference",
			source: "ref",
			want: ast.Reference{
				Token: token.Token{
					Type:  token.Identifier,
					Text:  "ref",
					Start: 0,
				},
			},
		},
		{
			name:   "function",
			source: "fn(in: Int) -> (Int) { return = in }",
			want: ast.Modification{
				Base: ast.FunctionType{
					FnToken: token.Token{
						Type:  token.Fn,
						Text:  "fn",
						Start: 0,
					},
					Inputs: []ast.Field{
						{
							Name: ast.Name{
								Token: token.Token{
									Type:  token.Identifier,
									Text:  "in",
									Start: 3,
								},
							},
							Type: ast.Reference{
								Token: token.Token{
									Type:  token.Identifier,
									Text:  "Int",
									Start: 7,
								},
							},
						},
					},
					Outputs: []ast.Field{
						{
							Type: ast.Reference{
								Token: token.Token{
									Type:  token.Identifier,
									Text:  "Int",
									Start: 16,
								},
							},
						},
					},
					CloseParen: token.Token{
						Type:  token.Character,
						Text:  ")",
						Start: 19,
					},
				},
				Block: ast.Block{
					LCurly: token.Token{
						Type:  token.Character,
						Text:  "{",
						Start: 21,
					},
					Assignments: []ast.Assignment{
						{
							Name: ast.Name{
								Token: token.Token{
									Type:  token.Identifier,
									Text:  "return",
									Start: 23,
								},
							},
							Expr: ast.Reference{
								Token: token.Token{
									Type:  token.Identifier,
									Text:  "in",
									Start: 32,
								},
							},
						},
					},
					RCurly: token.Token{
						Type:  token.Character,
						Text:  "}",
						Start: 35,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer(tc.source))
			got, err := ast.ParseExpr(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
