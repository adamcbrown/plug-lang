package ast

import (
	"fmt"
	"strconv"

	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/acbrown/plug-lang/types"
)

type Expr interface {
	Node
	AsType(*Context) types.Type
	Type(*Context) types.Type
	exprNode()
}

type ExprToken struct{}

func (ExprToken) exprNode() {}

func ParseExpr(p *parser.Parser) (Expr, *ParseErr) {
	return parsePostFixExpr(p)
}

func parsePostFixExpr(p *parser.Parser) (Expr, *ParseErr) {
	expr, err := parseAtomExpr(p)
	if err != nil {
		return nil, err
	}
	for {
		tok := p.ScanIgnoreWS()
		p.Unscan()
		if !p.InFunction() && tok.IsRune('{') {
			block, err := ParseBlock(p)
			if err != nil {
				return nil, err
			}
			expr = &Modification{
				Base:  expr,
				Block: block,
			}
			continue
		}
		return expr, nil
	}
}

func parseAtomExpr(p *parser.Parser) (Expr, *ParseErr) {
	tok := p.ScanIgnoreWS()
	switch tok.Type {
	case token.Integer:
		parsedInt, err := strconv.ParseInt(string(tok.Text), 0, 64)
		if err != nil {
			return nil, &ParseErr{
				Msg: err.Error(),
				Tok: tok,
			}
		}
		return &Constant[int]{
			Token: tok,
			Value: int(parsedInt),
		}, nil
	case token.Identifier:
		return &Reference{
			Token: tok,
		}, nil
	case token.Fn:
		p.Unscan()
		fnType, err := ParseFunctionType(p)
		if err != nil {
			return nil, err
		}
		return &fnType, nil
	default:
		return nil, &ParseErr{
			Msg: fmt.Sprintf("unknown token at expression. got %v", tok),
			Tok: tok,
		}
	}
}
