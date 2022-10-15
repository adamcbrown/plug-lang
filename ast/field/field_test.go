package field_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast/expr"
	"github.com/acbrown/plug-lang/ast/field"
	"github.com/acbrown/plug-lang/ast/name"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestField(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   field.Field
	}{
		{
			name:   "named field",
			source: "x: Int",
			want: field.Field{
				Name: name.Name{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  []rune("x"),
						Start: 0,
					},
				},
				Type: expr.Reference{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  []rune("Int"),
						Start: 3,
					},
				},
			},
		},
		{
			name:   "unnamed field",
			source: "Int",
			want: field.Field{
				Name: name.Name{},
				Type: expr.Reference{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  []rune("Int"),
						Start: 0,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer([]rune(tc.source)))
			got, err := field.Parse(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
