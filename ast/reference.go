package ast

import (
	"fmt"
	"sync"

	"github.com/acbrown/plug-lang/lexer/token"
	"github.com/acbrown/plug-lang/types"
)

type Reference struct {
	ExprToken
	Token token.Token

	typeOnce, asTypeOnce sync.Once
	typ, asType          types.Type
}

var _ Expr = &Reference{}

func (r *Reference) Start() int {
	return r.Token.StartPos()
}

func (r *Reference) End() int {
	return r.Token.EndPos()
}

func (r *Reference) AddReferences(ctx *Context) {
	name, ok := ctx.Resolve(r.Token.Text)
	if !ok {
		ctx.AddError(NodeErr{
			Msg: fmt.Sprintf("cannot resolve %s to node", r.Token.Text),
			N:   r,
		})
		return
	}
	ctx.AddReference(r, name)
}

func (r *Reference) Type(ctx *Context) types.Type {
	r.typeOnce.Do(func() {
		res, ok := ctx.Resolved(r)
		if !ok {
			ctx.AddError(NodeErr{
				Msg: "unable to resolve reference",
				N:   r,
			})
			r.typ = types.ErrorType
			return
		}
		r.typ = res.Type(ctx)
	})
	return r.typ
}

func (r *Reference) AsType(ctx *Context) types.Type {
	r.asTypeOnce.Do(func() {
		res, ok := ctx.Resolved(r)
		if !ok {
			ctx.AddError(NodeErr{
				Msg: "unable to resolve reference",
				N:   r,
			})
			r.typ = types.ErrorType
			return
		}
		r.typ = res.AsType(ctx)
	})
	return r.asType
}
