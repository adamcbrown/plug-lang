package ast

import (
	"testing"

	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestReferenceResolution(t *testing.T) {
	tcs := []struct {
		name string
		src  string
		want map[Reference]Expr
	}{
		{
			name: "basic",
			src:  "x = 1\ny = x\nz = y",
			want: map[Reference]Expr{
				{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "x",
						Start: 10,
					},
				}: &Constant[int]{
					Token: token.Token{
						Type:  token.Integer,
						Text:  "1",
						Start: 4,
					},
					Value: 1,
				},
				{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "y",
						Start: 16,
					},
				}: &Reference{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "x",
						Start: 10,
					},
				},
			},
		},
		{
			name: "recursive",
			src:  "x = x",
			want: map[Reference]Expr{
				{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "x",
						Start: 4,
					},
				}: &Reference{
					Token: token.Token{
						Type:  token.Identifier,
						Text:  "x",
						Start: 4,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer(tc.src))
			prgm, err := ParseProgram(p)
			if err != nil {
				t.Fatalf("ParseProgram(): err = %v", err)
			}

			ctx := NewContext()
			prgm.AddReferences(ctx)

			if len(ctx.errors) > 0 {
				t.Fatalf("ReferenceResolution(): err = %v", ctx.errors)
			}

			unPtrRefs := make(map[Reference]Expr, len(ctx.resolutions))
			for ref, expr := range ctx.resolutions {
				unPtrRefs[*ref] = expr
			}

			if diff := cmp.Diff(tc.want, unPtrRefs, IgnoreOpt); diff != "" {
				t.Errorf("ReferenceResolution() got diff (-want, +got):\n%s", diff)
			}
		})
	}
}
