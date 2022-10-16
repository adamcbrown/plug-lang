package ast_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestFunction(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   ast.FunctionType
	}{
		{
			name:   "unnamed fn output",
			source: "fn() -> (Int)",
			want: ast.FunctionType{
				FnToken: token.Token{
					Type:  token.Fn,
					Text:  []rune("fn"),
					Start: 0,
				},
				Inputs: nil,
				Outputs: []ast.Field{
					{
						Name: ast.Name{},
						Type: ast.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 9,
							},
						},
					},
				},
				CloseParen: token.Token{
					Type:  token.Character,
					Text:  []rune(")"),
					Start: 12,
				},
			},
		},
		{
			name:   "named fn output",
			source: "fn() -> (out: Int)",
			want: ast.FunctionType{
				FnToken: token.Token{
					Type:  token.Fn,
					Text:  []rune("fn"),
					Start: 0,
				},
				Inputs: nil,
				Outputs: []ast.Field{
					{
						Name: ast.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("out"),
								Start: 9,
							},
						},
						Type: ast.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 14,
							},
						},
					},
				},
				CloseParen: token.Token{
					Type:  token.Character,
					Text:  []rune(")"),
					Start: 17,
				},
			},
		},
		{
			name:   "fn with input",
			source: "fn(x: Int) -> (y: Int)",
			want: ast.FunctionType{
				FnToken: token.Token{
					Type:  token.Fn,
					Text:  []rune("fn"),
					Start: 0,
				},
				Inputs: []ast.Field{
					{
						Name: ast.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("x"),
								Start: 3,
							},
						},
						Type: ast.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 6,
							},
						},
					},
				},
				Outputs: []ast.Field{
					{
						Name: ast.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("y"),
								Start: 15,
							},
						},
						Type: ast.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 18,
							},
						},
					},
				},
				CloseParen: token.Token{
					Type:  token.Character,
					Text:  []rune(")"),
					Start: 21,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer([]rune(tc.source)))
			got, err := ast.ParseFunctionType(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
