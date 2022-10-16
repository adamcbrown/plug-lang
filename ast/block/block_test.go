package block_test

import (
	"testing"

	"github.com/acbrown/plug-lang/ast/assignment"
	"github.com/acbrown/plug-lang/ast/block"
	"github.com/acbrown/plug-lang/ast/expr"
	"github.com/acbrown/plug-lang/ast/name"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/google/go-cmp/cmp"
)

func TestBlock(t *testing.T) {
	tcs := []struct {
		name   string
		source string
		want   block.Block
	}{
		{
			name:   "empty block",
			source: "{}",
			want: block.Block{
				LCurly: token.Token{
					Type:  token.Character,
					Text:  []rune("{"),
					Start: 0,
				},
				RCurly: token.Token{
					Type:  token.Character,
					Text:  []rune("}"),
					Start: 1,
				},
			},
		},
		{
			name:   "block with elements",
			source: "{x = 1}",
			want: block.Block{
				LCurly: token.Token{
					Type:  token.Character,
					Text:  []rune("{"),
					Start: 0,
				},
				Assignments: []assignment.Assignment{
					{
						Name: name.Name{
							Token: token.Token{
								Type:  token.Identifier,
								Text:  []rune("x"),
								Start: 1,
							},
						},
						Expr: expr.Constant[int]{
							Token: token.Token{
								Type:  token.Integer,
								Text:  []rune("1"),
								Start: 5,
							},
							Value: 1,
						},
					},
				},
				RCurly: token.Token{
					Type:  token.Character,
					Text:  []rune("}"),
					Start: 6,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser(lexer.NewLexer([]rune(tc.source)))
			got, err := block.Parse(p)
			if err != nil {
				t.Fatalf("Parse(): err = %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
