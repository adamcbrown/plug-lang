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
					Text:  []rune("10"),
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
					Text:  []rune("ref"),
					Start: 0,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer([]rune(tc.source)))
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
