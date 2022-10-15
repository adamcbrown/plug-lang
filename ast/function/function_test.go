package function_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast/expr"
	"github.com/acbrown/plug-lang/ast/field"
	"github.com/acbrown/plug-lang/ast/function"
	"github.com/acbrown/plug-lang/ast/name"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestFunction(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   function.FunctionType
	}{
		{
			name:   "unnamed fn output",
			source: "fn() -> Int",
			want: function.FunctionType{
				FnToken: token.Token{
					Type:  token.Fn,
					Text:  []rune("fn"),
					Start: 0,
				},
				Inputs: nil,
				Outputs: []field.Field{
					{
						Name: name.Name{},
						Type: expr.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 8,
							},
						},
					},
				},
			},
		},
		{
			name:   "named fn output with paren",
			source: "fn() -> (out: Int)",
			want: function.FunctionType{
				FnToken: token.Token{
					Type:  token.Fn,
					Text:  []rune("fn"),
					Start: 0,
				},
				Inputs: nil,
				Outputs: []field.Field{
					{
						Name: name.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("out"),
								Start: 9,
							},
						},
						Type: expr.Reference{
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
			name:   "named fn output no paren",
			source: "fn() -> out: Int",
			want: function.FunctionType{
				FnToken: token.Token{
					Type:  token.Fn,
					Text:  []rune("fn"),
					Start: 0,
				},
				Inputs: nil,
				Outputs: []field.Field{
					{
						Name: name.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("out"),
								Start: 8,
							},
						},
						Type: expr.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 13,
							},
						},
					},
				},
			},
		},
		{
			name:   "fn with input",
			source: "fn(x: Int) -> y: Int",
			want: function.FunctionType{
				FnToken: token.Token{
					Type:  token.Fn,
					Text:  []rune("fn"),
					Start: 0,
				},
				Inputs: []field.Field{
					{
						Name: name.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("x"),
								Start: 3,
							},
						},
						Type: expr.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 6,
							},
						},
					},
				},
				Outputs: []field.Field{
					{
						Name: name.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("y"),
								Start: 14,
							},
						},
						Type: expr.Reference{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("Int"),
								Start: 17,
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer([]rune(tc.source)))
			got, err := function.Parse(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
