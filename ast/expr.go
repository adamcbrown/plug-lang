package ast

import (
	"fmt"
	"strconv"

	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

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
		if tok.IsRune('{') {
			p.Unscan()
			block, err := ParseBlock(p)
			if err != nil {
				return nil, err
			}
			expr = Modification{
				Base:  expr,
				Block: block,
			}
		}
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
		return ParseFunctionType(p)
	default:
		return nil, &ParseErr{
			Msg: fmt.Sprintf("unknown token at expression. got %v", tok),
			Tok: tok,
		}
	}
}
