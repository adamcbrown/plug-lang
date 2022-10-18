package ast

import (
	"sync"

	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/parser"
	"github.com/acbrown/plug-lang/types"
)

type FunctionType struct {
	ExprToken
	FnToken    token.Token
	Inputs     []Field
	Outputs    []Field
	CloseParen token.Token

	typeOnce sync.Once
	typ      types.FunctionType
}

var _ Expr = &FunctionType{}

func (ft *FunctionType) Start() int {
	return ft.FnToken.StartPos()
}

func (ft *FunctionType) End() int {
	if ft.CloseParen.IsValid() {
		return ft.CloseParen.EndPos()
	}
	return ft.Outputs[0].End()
}

func (ft *FunctionType) AddReferences(ctx *Context) {}

func (ft *FunctionType) Type(ctx *Context) types.Type {
	return types.TypeType
}

func (ft *FunctionType) AsType(ctx *Context) types.Type {
	ft.typeOnce.Do(func() {
		inputs := make([]types.Type, 0, len(ft.Inputs))
		for _, input := range ft.Inputs {
			inputs = append(inputs, input.Type.Type(ctx))
		}

		outputs := make([]types.Type, 0, len(ft.Outputs))
		for _, output := range ft.Outputs {
			outputs = append(outputs, output.Type.Type(ctx))
		}

		ft.typ = types.FunctionType{
			Inputs:  inputs,
			Outputs: outputs,
		}
	})
	return ft.typ
}

func ParseFunctionType(p *parser.Parser) (FunctionType, *ParseErr) {
	p.EnterType()
	defer p.ExitType()

	fn := p.ScanIgnoreWS()
	if fn.Type != token.Fn {
		return FunctionType{}, &ParseErr{
			Msg: "expected `fn` token at start of function",
			Tok: fn,
		}
	}

	if tok := p.ScanIgnoreWS(); !tok.IsRune('(') {
		return FunctionType{}, &ParseErr{
			Msg: "expected `(` token after `fn`",
			Tok: tok,
		}
	}

	var inputs []Field
	for {
		tok := p.ScanIgnoreWS()
		if tok.IsRune(')') {
			break
		}
		p.Unscan()

		f, err := ParseField(p)
		if err != nil {
			return FunctionType{}, err
		}
		inputs = append(inputs, f)

		if tok := p.Scan(); !tok.IsRune(',') && !tok.IsRune(')') {
			return FunctionType{}, &ParseErr{
				Msg: "expected `,` or `)` after field in function outputs",
				Tok: tok,
			}
		} else if tok.IsRune(')') {
			p.Unscan()
		}
	}

	arrowStart := p.ScanIgnoreWS()
	if !arrowStart.IsRune('-') {
		return FunctionType{}, &ParseErr{
			Msg: "expected `->` token after fn arguments",
			Tok: arrowStart,
		}
	}
	if tok := p.Scan(); !tok.IsRune('>') {
		return FunctionType{}, &ParseErr{
			Msg: "expected `->` token after fn arguments",
			Tok: arrowStart,
		}
	}

	if tok := p.ScanIgnoreWS(); tok.IsRune('(') {
		var outputs []Field
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

			f, err := ParseField(p)
			if err != nil {
				return FunctionType{}, err
			}
			outputs = append(outputs, f)

			if tok := p.Scan(); !tok.IsRune(',') && !tok.IsRune(')') {
				return FunctionType{}, &ParseErr{
					Msg: "expected `,` or `)` after field in function outputs",
					Tok: tok,
				}
			} else if tok.IsRune(')') {
				p.Unscan()
			}
		}
	}

	p.Unscan()
	output, err := ParseField(p)
	if err != nil {
		return FunctionType{}, err
	}

	return FunctionType{
		FnToken:    fn,
		Inputs:     inputs,
		Outputs:    []Field{output},
		CloseParen: token.Token{},
	}, nil
}
