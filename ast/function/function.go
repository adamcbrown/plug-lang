package function

import (
	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/ast/expr"
	"github.com/acbrown/plug-lang/ast/field"
	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
)

type FunctionType struct {
	expr.ExprToken
	FnToken    token.Token
	Inputs     []field.Field
	Outputs    []field.Field
	CloseParen token.Token
}

func (ft FunctionType) Start() int {
	return ft.FnToken.StartPos()
}

func (ft FunctionType) End() int {
	if ft.CloseParen.IsValid() {
		return ft.CloseParen.EndPos()
	}
	return ft.Outputs[0].End()
}

func Parse(p *parser.Parser) (FunctionType, *ast.ParseErr) {
	fn := p.ScanIgnoreWS()
	if fn.Type != token.Fn {
		return FunctionType{}, &ast.ParseErr{
			Msg: "expected `fn` token at start of function",
			Tok: fn,
		}
	}

	if tok := p.ScanIgnoreWS(); !tok.IsRune('(') {
		return FunctionType{}, &ast.ParseErr{
			Msg: "expected `(` token after `fn`",
			Tok: tok,
		}
	}

	var inputs []field.Field
	for {
		tok := p.ScanIgnoreWS()
		if tok.IsRune(')') {
			break
		}
		p.Unscan()

		f, err := field.Parse(p)
		if err != nil {
			return FunctionType{}, err
		}
		inputs = append(inputs, f)

		if tok := p.Scan(); !tok.IsRune(',') && !tok.IsRune(')') {
			return FunctionType{}, &ast.ParseErr{
				Msg: "expected `,` or `)` after field in function outputs",
				Tok: tok,
			}
		} else if tok.IsRune(')') {
			p.Unscan()
		}
	}

	arrowStart := p.ScanIgnoreWS()
	if !arrowStart.IsRune('-') {
		return FunctionType{}, &ast.ParseErr{
			Msg: "expected `->` token after fn arguments",
			Tok: arrowStart,
		}
	}
	if tok := p.Scan(); !tok.IsRune('>') {
		return FunctionType{}, &ast.ParseErr{
			Msg: "expected `->` token after fn arguments",
			Tok: arrowStart,
		}
	}

	if tok := p.ScanIgnoreWS(); tok.IsRune('(') {
		var outputs []field.Field
		for {
			tok := p.ScanIgnoreWS()
			if tok.IsRune(')') {
				return FunctionType{
					FnToken:    fn,
					Inputs:     inputs,
					Outputs:    outputs,
					CloseParen: tok,
				}, nil
			}
			p.Unscan()

			f, err := field.Parse(p)
			if err != nil {
				return FunctionType{}, err
			}
			outputs = append(inputs, f)

			if tok := p.Scan(); !tok.IsRune(',') && !tok.IsRune(')') {
				return FunctionType{}, &ast.ParseErr{
					Msg: "expected `,` or `)` after field in function outputs",
					Tok: tok,
				}
			} else if tok.IsRune(')') {
				p.Unscan()
			}
		}
	}

	p.Unscan()
	output, err := field.Parse(p)
	if err != nil {
		return FunctionType{}, err
	}

	return FunctionType{
		FnToken:    fn,
		Inputs:     inputs,
		Outputs:    []field.Field{output},
		CloseParen: token.Token{},
	}, nil
}
