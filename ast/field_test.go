package ast_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestField(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   ast.Field
	}{
		{
			name:   "named field",
			source: "x: Int",
			want: ast.Field{
				Name: ast.Name{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "x",
						Start: 0,
					},
				},
				Type: &ast.Reference{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "Int",
						Start: 3,
					},
				},
			},
		},
		{
			name:   "unnamed field",
			source: "Int",
			want: ast.Field{
				Name: ast.Name{},
				Type: &ast.Reference{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "Int",
						Start: 0,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer(tc.source))
			got, err := ast.ParseField(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
