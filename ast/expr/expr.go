package expr

import (
	"fmt"
	"strconv"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/ast/function"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type Expr interface {
	ast.Node
	exprNode()
}

type ExprToken struct{}

func (ExprToken) exprNode() {}

func Parse(p *parser.Parser) (Expr, *ast.ParseErr) {
	return parseAtomExpr(p)
}

func parseAtomExpr(p *parser.Parser) (Expr, *ast.ParseErr) {
	tok := p.ScanIgnoreWS()
	switch tok.Type {
	case token.Integer:
		parsedInt, err := strconv.ParseInt(string(tok.Text), 0, 64)
		if err != nil {
			return nil, &ast.ParseErr{
				Msg: err.Error(),
				Tok: tok,
			}
		}
		return Constant[int]{
			Token: tok,
			Value: int(parsedInt),
		}, nil
	case token.Identifier:
		return Reference{
			Token: tok,
		}, nil
	case token.Fn:
		p.Unscan()
		return function.Parse(p)
	default:
		return nil, &ast.ParseErr{
			Msg: fmt.Sprintf("unknown token at expression. got %v", tok),
			Tok: tok,
		}
	}
}
